package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "do-data-go/migrations"
)

type ExportRecord struct {
	Value            float64
	Domainname       string
	Typename         string
	Organizationtype string
	Year             int
}

func GetOrganisationId(app *pocketbase.PocketBase, event *core.RequestEvent) string {
	organizationId := event.Auth.Get("organization")
	return organizationId.(string)
}

func GetDomainId(app core.App, domainName string) (string, error) {
	domains, err := app.FindAllRecords("domains")
	if err != nil {
		return "", nil
	}
	index := slices.IndexFunc(domains, func(record *core.Record) bool { return record.Get("name").(string) == domainName })
	if index > -1 {
		return domains[index].Id, nil
	} else {
		return "", fmt.Errorf("Could not find: %s", domainName)
	}
}

func GetAggregationId(app core.App, aggregationName string) (string, error) {
	aggregations, err := app.FindAllRecords("time_aggregations")
	if err != nil {
		return "", nil
	}
	index := slices.IndexFunc(aggregations, func(record *core.Record) bool { return record.Get("name").(string) == aggregationName })
	if index > -1 {
		return aggregations[index].Id, nil
	} else {
		return "", fmt.Errorf("Could not find: %s", aggregationName)
	}
}

func GetEntryTypeId(app core.App, typeName string) (string, error) {
	entryTypes, err := app.FindAllRecords("types")
	if err != nil {
		return "", nil
	}
	index := slices.IndexFunc(entryTypes, func(record *core.Record) bool { return record.Get("name").(string) == typeName })
	if index > -1 {
		return entryTypes[index].Id, nil
	} else {
		return "", fmt.Errorf("Could not find: %s", typeName)
	}
}

func GetUnitId(app core.App, unitName string) (string, error) {
	unitTypes, err := app.FindAllRecords("units")
	if err != nil {
		return "", nil
	}
	index := slices.IndexFunc(unitTypes, func(record *core.Record) bool { return record.Get("name").(string) == unitName })
	if index > -1 {
		return unitTypes[index].Id, nil
	} else {
		return "", fmt.Errorf("Could not find: %s", unitName)
	}
}

func exportEntries(app *pocketbase.PocketBase, year int) ([]ExportRecord, error) {
	entries := []ExportRecord{}
	err := app.DB().
		Select(
			"SUM(view_entries.value) as value",
			"view_entries.domainName as domainname",
			"view_entries.typeName as typename",
			"organization_categories.name as organizationtype",
			"view_entries.year as year",
		).
		From("view_entries").
		LeftJoin("organizations", dbx.NewExp("view_entries.organization = organizations.id")).
		LeftJoin("organization_categories", dbx.NewExp("organizations.category = organization_categories.id")).
		Where(dbx.HashExp{"year": fmt.Sprintf("%d", year)}).
		GroupBy("domainname", "typename", "organizationtype", "year").
		All(&entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

type ExportCheck struct {
	Count int
}

func canExport(app pocketbase.PocketBase) bool {
	exportCheck := ExportCheck{}
	app.DB().
		Select("COUNT(DISTINCT organization) as count").
		From("entries").
		One(&exportCheck)

	organizationsWithEntries := exportCheck.Count

	export := os.Getenv("MUNICIPALITY_EXPORT")
	if len(export) == 0 {
		return false
	}
	exportCount, err := strconv.Atoi(export)
	if err != nil {
		return false
	}
	if organizationsWithEntries >= exportCount {
		return true
	}
	return false
}

func main() {
	app := pocketbase.New()

	redashAPIKey := os.Getenv("REDASH_API_KEY")
	redashBaseURL := os.Getenv("REDASH_BASE_URL")

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun, /* Update migrations when changes are made using the dashboard */
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Register export check endpoint
		se.Router.GET("/api/exportcheck", func(e *core.RequestEvent) error {
			return e.JSON(200, map[string]bool{"canExport": canExport(*app)})
		}).Bind(apis.RequireAuth())

		// Register export endpoint
		se.Router.GET("/api/export/{year}", func(e *core.RequestEvent) error {

			if !canExport(*app) {
				return e.ForbiddenError("Export not allowed due to insufficient parcipating organizations", nil)
			}

			yearParam := e.Request.PathValue("year")
			year, err := strconv.Atoi(yearParam)
			if err != nil {
				return e.BadRequestError("Year is not the correct format", nil)
			}

			entries, err := exportEntries(app, year)
			if err != nil {
				app.Logger().Debug(err.Error())
				return e.InternalServerError("Could not export data", nil)
			}

			printer := message.NewPrinter(language.Danish)
			var csvBuilder bytes.Buffer
			csvBuilder.WriteString("Value;Domain;Type;OragnizationType;Year\n")
			for _, entry := range entries {
				csvString := printer.Sprintf("%.2f;%s;%s;%s;%#d\n", entry.Value, entry.Domainname, entry.Typename, entry.Organizationtype, entry.Year)
				csvBuilder.WriteString(csvString)
			}

			e.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=dodata-%s.csv", yearParam))
			return e.Blob(200, "text/csv", csvBuilder.Bytes())
		}).Bind(apis.RequireAuth())

		// register "GET /api/visualization" route (allowed only for authenticated users)
		se.Router.GET("/api/visualization", func(e *core.RequestEvent) error {
			queryId := e.Request.URL.Query().Get("queryId")
			visualizationId := e.Request.URL.Query().Get("visualizationId")
			embedURL := redashBaseURL + "/embed/query/" + queryId + "/visualization/" + visualizationId + "?api_key=" + redashAPIKey
			return e.JSON(http.StatusOK, map[string]string{"embedURL": embedURL})
		}).Bind(apis.RequireAuth())

		// Endpoint for upload .xlsx files
		se.Router.POST("/api/upload", func(e *core.RequestEvent) error {
			/* Get file from request */
			files, err := e.FindUploadedFiles("file")
			if err != nil {
				return e.BadRequestError("Could not retrieve file from request", nil)
			}
			xslxFile := files[0]

			/* Create new file_imports record with request file */
			fileImports, err := app.FindCollectionByNameOrId("file_imports")
			if err != nil {
				return e.InternalServerError(err.Error(), nil)
			}
			fileImport := core.NewRecord(fileImports)
			fileImport.Set("file", files[0])
			saveErr := app.Save(fileImport)
			if saveErr != nil {
				return e.InternalServerError(saveErr.Error(), nil)
			}

			/* Process file contents */
			energyDataExtractor := energyDataExtractor{app: app}
			energyDataExtractor.Initialize(xslxFile)
			energyData, err := energyDataExtractor.GetEnergyData()
			if err != nil {
				if _, ok := err.(TimeSpanError); ok {
					return e.JSON(http.StatusBadRequest, "year_error")
				}
				app.Logger().Debug(err.Error())
			}

			/* Create energy entries */
			organizationId := GetOrganisationId(app, e)
			err = createEnergyDataRecords(app, &energyData, fileImport, organizationId)
			if err != nil {
				app.Logger().Error(err.Error())
				return e.InternalServerError(err.Error(), nil)
			}

			err = deleteRecommendationStatistics(app, organizationId)
			if err != nil {
				app.Logger().Error(err.Error())
				return e.InternalServerError(err.Error(), nil)
			}

			err = ProcessRules(app)
			if err != nil {
				app.Logger().Debug(err.Error())
			}

			return e.NoContent(http.StatusOK)
		}).Bind(apis.RequireAuth())

		// register "DELETE /api/dataentry/{id}" route (allowed only for authenticated users)
		se.Router.DELETE("/api/dataentries/{year}", func(e *core.RequestEvent) error {
			// Get year from request
			year := e.Request.PathValue("year")
			if year == "" {
				return e.BadRequestError("Year is required", nil)
			}

			// Get the authenticated user's organization
			organizationId := GetOrganisationId(app, e)

			// Get entries to delete, based on organization and year
			entries, err := app.FindAllRecords("entries",
				dbx.HashExp{"organization": organizationId},
				dbx.NewExp("strftime('%Y', timestamp) = {:year}", dbx.Params{"year": year}),
			)
			if err != nil {
				app.Logger().Debug(err.Error())
				return e.JSON(http.StatusInternalServerError, nil)
			}
			app.Logger().Debug(fmt.Sprintf("Entries: %v", entries))

			// Delete all entries in a transaction so no dangling entries a deleted year can be present
			err = app.RunInTransaction(func(txApp core.App) error {
				for _, entry := range entries {
					err := txApp.Delete(entry)
					if err != nil {
						return err
					}
				}
				return nil
			})
			if err != nil {
				app.Logger().Debug(err.Error())
			}
			return e.JSON(http.StatusOK, map[string]string{"message": "Data entry deleted successfully"})
		}).Bind(apis.RequireAuth())

		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

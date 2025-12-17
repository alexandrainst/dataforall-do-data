package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {

		organizationCategories := []string{
			"Servicevirksomhed",
			"Produktionsvirksomhed",
			"Handelsvirksomhed",
		}
		createNamedCollectionRecords(app, "organization_categories", organizationCategories)
		
		situations := []string{
			"trending_up",
			"trending_down",
			"trend_steady",
			"above",
			"below",
			"average",
		}
		createNamedCollectionRecords(app, "situations", situations)

		return nil

	}, func(app core.App) error {

		organizationCategories, err := app.FindCollectionByNameOrId("organization_categories")
		if err != nil {
			return err
		}
		err = app.TruncateCollection(organizationCategories)
		if err != nil {
			return err
		}

		situations, err := app.FindCollectionByNameOrId("situations")
		if err != nil {
			return err
		}
		err = app.TruncateCollection(situations)
		if err != nil {
			return err
		}

		return nil
	})
}

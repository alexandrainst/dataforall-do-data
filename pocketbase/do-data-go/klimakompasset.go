package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/xuri/excelize/v2"
)

type energyDataExtractor struct {
	app  *pocketbase.PocketBase
	file *excelize.File
}

type energyConsumption struct {
	nonRenewable float64
	renewable    float64
	unit         string
}

type energyConsumptions struct {
	electricity *energyConsumption /* Will be nil if no electricity data can be extracted from the file */
	heat        *energyConsumption /* Will be nil if no heat data can be extracted from the file */
	fuel        *energyConsumption /* Will be nil if no fuel data can be extracted from the file */
}

type co2Emission struct {
	value float64
	unit  string
}

type co2Emissions struct {
	electricity *co2Emission /* Will be nil if no electricity data can be extracted from the file */
	heat        *co2Emission /* Will be nil if no heat data can be extracted from the file */
	fuel        *co2Emission /* Will be nil if no fuel data can be extracted from the file */
}

type energyData struct {
	startDate   time.Time
	endDate     time.Time
	consumption energyConsumptions
	emission    co2Emissions
}

type TimeSpanError struct{}

func (e TimeSpanError) Error() string {
	return "Data does not cover the whole year"
}

func getCellValue[T any](e *energyDataExtractor, sheet string, cell string, convert func(string) (T, error)) (T, error) {
	if e.file != nil {
		cellValue, err := e.file.GetCellValue(sheet, cell)
		e.app.Logger().Debug(cellValue)
		if err == nil {
			return convert(cellValue)
		}
	}
	return *new(T), fmt.Errorf("Could not get cell from worksheet: %s, cell: %s", sheet, cell)
}

func (e *energyDataExtractor) Initialize(file *filesystem.File) {
	reader, err := file.Reader.Open()
	if err != nil {
		e.app.Logger().Debug(err.Error())
	}
	f, ferr := excelize.OpenReader(reader)
	if ferr != nil {
		e.app.Logger().Debug(ferr.Error())
	}
	e.file = f
}

func convertToTime(value string) (time.Time, error) {
	return time.Parse("01-02-06", value)
}

func convertToFloat(value string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(value, ",", ""), 64)
}

func converToUnitString(value string) (string, error) {
	return value, nil // TODO: Do something here?
}

// TODO: Consider at better function signature, use struct?
func (e *energyDataExtractor) GetEnergyConsumption() (energyConsumptions, bool, bool, bool, error) {
	energyConsumptions := energyConsumptions{}

	/* Extract electricity */
	skipElectricity := false
	electricityRenewable, err := getCellValue(e, "E-nøgletal (ESG)", "D35", convertToFloat)
	if err != nil {
		/* If the value cannot be parsed skip all related values */
		if errors.Is(err, strconv.ErrSyntax) {
			skipElectricity = true
		} else {
			return energyConsumptions, true, true, true, err
		}
	}
	if !skipElectricity {
		electricityNonRenewable, err := getCellValue(e, "E-nøgletal (ESG)", "E35", convertToFloat)
		if err != nil {
			return energyConsumptions, true, true, true, err
		}
		electricityUnit, err := getCellValue(e, "E-nøgletal (ESG)", "G35", converToUnitString)
		if err != nil {
			return energyConsumptions, true, true, true, err
		}
		electricityConsumption := energyConsumption{
			renewable:    electricityRenewable,
			nonRenewable: electricityNonRenewable,
			unit:         electricityUnit,
		}
		energyConsumptions.electricity = &electricityConsumption
	}

	/* Extract fuel */
	skipFuel := false
	fuelRenewable, err := getCellValue(e, "E-nøgletal (ESG)", "D36", convertToFloat)
	if err != nil {
		/* If the value cannot be parsed skip all related values */
		if errors.Is(err, strconv.ErrSyntax) {
			skipFuel = true
		} else {
			return energyConsumptions, true, true, true, err
		}
	}
	if !skipFuel {
		fuelNonRenewable, err := getCellValue(e, "E-nøgletal (ESG)", "E36", convertToFloat)
		if err != nil {
			return energyConsumptions, true, true, true, err
		}
		fuelUnit, err := getCellValue(e, "E-nøgletal (ESG)", "G36", converToUnitString)
		if err != nil {
			return energyConsumptions, true, true, true, err
		}
		fuelConsumption := energyConsumption{
			renewable:    fuelRenewable,
			nonRenewable: fuelNonRenewable,
			unit:         fuelUnit,
		}
		energyConsumptions.fuel = &fuelConsumption
	}

	/* Extract heat */
	skipHeat := false
	heatRenewable, err := getCellValue(e, "E-nøgletal (ESG)", "D37", convertToFloat)
	if err != nil {
		/* If the value cannot be parsed skip all related values */
		if errors.Is(err, strconv.ErrSyntax) {
			skipHeat = true
		} else {
			return energyConsumptions, true, true, true, err
		}
	}
	if !skipHeat {
		heatNonRenewable, err := getCellValue(e, "E-nøgletal (ESG)", "E37", convertToFloat)
		if err != nil {
			return energyConsumptions, true, true, true, err
		}
		heatUnit, err := getCellValue(e, "E-nøgletal (ESG)", "G37", converToUnitString)
		if err != nil {
			return energyConsumptions, true, true, true, err
		}
		heatConsumption := energyConsumption{
			renewable:    heatRenewable,
			nonRenewable: heatNonRenewable,
			unit:         heatUnit,
		}
		energyConsumptions.heat = &heatConsumption
	}

	return energyConsumptions, skipElectricity, skipFuel, skipHeat, nil
}

func (e *energyDataExtractor) GetCO2Emissions(skipElectricity bool, skipFuel bool, skipHeat bool) (co2Emissions, error) {
	co2Emissions := co2Emissions{}

	e.app.Logger().Debug(fmt.Sprintf("Skip electricity: %t", skipElectricity))
	if !skipElectricity {
		electricityScope1, err := getCellValue(e, "Delresultater (GHG)", "G25", convertToFloat)
		if err != nil {
			return co2Emissions, err
		}
		electricityEmission := co2Emission{
			value: electricityScope1,
			unit:  "Ton",
		}
		co2Emissions.electricity = &electricityEmission
	}

	e.app.Logger().Debug(fmt.Sprintf("Skip fuel: %t", skipFuel))
	if !skipFuel {
		fuelScope1, err := getCellValue(e, "Delresultater (GHG)", "G30", convertToFloat)
		if err != nil {
			return co2Emissions, err
		}
		fuelEmission := co2Emission{
			value: fuelScope1,
			unit:  "Ton",
		}
		co2Emissions.fuel = &fuelEmission
	}

	e.app.Logger().Debug(fmt.Sprintf("Skip heat: %t", skipHeat))
	if !skipHeat {
		heatScope1, err := getCellValue(e, "Delresultater (GHG)", "G29", convertToFloat)
		if err != nil {
			return co2Emissions, err
		}
		heatEmission := co2Emission{
			value: heatScope1,
			unit:  "Ton",
		}
		co2Emissions.heat = &heatEmission
	}

	return co2Emissions, nil
}

func (e *energyDataExtractor) GetEnergyData() (energyData, error) {
	energyData := energyData{}
	if e.file == nil {
		return energyData, errors.New("Imported file not initialized")
	}

	/* Extract start date */
	startDate, err := getCellValue(e, "Stamdata", "B14", convertToTime)
	if err != nil {
		return energyData, err
	}
	energyData.startDate = startDate

	/* Extract end date */
	endDate, err := getCellValue(e, "Stamdata", "B17", convertToTime)
	if err != nil {
		return energyData, err
	}
	energyData.endDate = endDate

	/* Check if file covers whole year, if not return error */
	startYear, startMonth, startDay := startDate.Date()
	endYear, endMonth, endDay := endDate.Date()
	isWholeYear := startYear == endYear && startMonth == 1 && startDay == 1 && endMonth == 12 && endDay == 31
	if !isWholeYear {
		return energyData, TimeSpanError{}
	}

	/* Extract energy consumption */
	energyConsumption, skipElectricity, skipFuel, skipHeat, err := e.GetEnergyConsumption()
	if err != nil {
		return energyData, err
	}
	energyData.consumption = energyConsumption

	/* Extract co2 emissions */
	co2Emissions, err := e.GetCO2Emissions(skipElectricity, skipFuel, skipHeat)
	if err != nil {
		return energyData, err
	}
	energyData.emission = co2Emissions

	return energyData, nil
}

/*
 * Create all records related to klimakompasset data import
 */
func createEnergyDataRecords(app *pocketbase.PocketBase, energyData *energyData, fileImport *core.Record, organizationId string) error {
	/* Create all records in transaction */
	createEnergyDataRecordsError := app.RunInTransaction(func(transactionApp core.App) error {
		entries, err := transactionApp.FindCollectionByNameOrId("entries")
		if err != nil {
			return err
		}
		renewableTypeId, err := GetEntryTypeId(transactionApp, "Energy Consumption Renewable")
		if err != nil {
			return err
		}
		nonRenewableTypeId, err := GetEntryTypeId(transactionApp, "Energy Consumption Non-renewable")
		if err != nil {
			return err
		}
		co2EmissionTypeId, err := GetEntryTypeId(transactionApp, "CO2 Emission")
		if err != nil {
			return err
		}
		aggregationYearId, err := GetAggregationId(transactionApp, "Year")
		if err != nil {
			return err
		}
		electricityDomain, err := GetDomainId(transactionApp, "El")
		if err != nil {
			return err
		}
		fuelDomain, err := GetDomainId(transactionApp, "Gas")
		if err != nil {
			return err
		}
		heatDomain, err := GetDomainId(transactionApp, "Varme")
		if err != nil {
			return err
		}

		if energyData.consumption.electricity != nil {
			unitId, err := GetUnitId(transactionApp, energyData.consumption.electricity.unit)
			if err != nil {
				return err
			}

			electricityRenewableConsumptionEntry := core.NewRecord(entries)
			electricityRenewableConsumptionEntry.Set("organization", organizationId)
			electricityRenewableConsumptionEntry.Set("domain", electricityDomain)
			electricityRenewableConsumptionEntry.Set("type", renewableTypeId)
			electricityRenewableConsumptionEntry.Set("aggregation", aggregationYearId)
			electricityRenewableConsumptionEntry.Set("timestamp", energyData.startDate)
			electricityRenewableConsumptionEntry.Set("unit", unitId)
			electricityRenewableConsumptionEntry.Set("value", energyData.consumption.electricity.renewable)
			saveErr := transactionApp.Save(electricityRenewableConsumptionEntry)
			if saveErr != nil {
				return saveErr
			}

			electricityNonRenewableConsumptionEntry := core.NewRecord(entries)
			electricityNonRenewableConsumptionEntry.Set("organization", organizationId)
			electricityNonRenewableConsumptionEntry.Set("domain", electricityDomain)
			electricityNonRenewableConsumptionEntry.Set("type", nonRenewableTypeId)
			electricityNonRenewableConsumptionEntry.Set("aggregation", aggregationYearId)
			electricityNonRenewableConsumptionEntry.Set("timestamp", energyData.startDate)
			electricityNonRenewableConsumptionEntry.Set("unit", unitId)
			electricityNonRenewableConsumptionEntry.Set("value", energyData.consumption.electricity.nonRenewable)
			saveErr = transactionApp.Save(electricityNonRenewableConsumptionEntry)
			if saveErr != nil {
				return saveErr
			}
			fileImport.Set("entires+", electricityRenewableConsumptionEntry)
			transactionApp.Save(fileImport)
		}

		if energyData.consumption.fuel != nil {
			unitId, err := GetUnitId(transactionApp, energyData.consumption.fuel.unit)
			if err != nil {
				return err
			}

			fuelRenewableConsumptionEntry := core.NewRecord(entries)
			fuelRenewableConsumptionEntry.Set("organization", organizationId)
			fuelRenewableConsumptionEntry.Set("domain", fuelDomain)
			fuelRenewableConsumptionEntry.Set("type", renewableTypeId)
			fuelRenewableConsumptionEntry.Set("aggregation", aggregationYearId)
			fuelRenewableConsumptionEntry.Set("timestamp", energyData.startDate)
			fuelRenewableConsumptionEntry.Set("unit", unitId)
			fuelRenewableConsumptionEntry.Set("value", energyData.consumption.fuel.renewable)
			saveErr := transactionApp.Save(fuelRenewableConsumptionEntry)
			if saveErr != nil {
				return saveErr
			}

			fuelNonRenewableConsumptionEntry := core.NewRecord(entries)
			fuelNonRenewableConsumptionEntry.Set("organization", organizationId)
			fuelNonRenewableConsumptionEntry.Set("domain", fuelDomain)
			fuelNonRenewableConsumptionEntry.Set("type", nonRenewableTypeId)
			fuelNonRenewableConsumptionEntry.Set("aggregation", aggregationYearId)
			fuelNonRenewableConsumptionEntry.Set("timestamp", energyData.startDate)
			fuelNonRenewableConsumptionEntry.Set("unit", unitId)
			fuelNonRenewableConsumptionEntry.Set("value", energyData.consumption.fuel.nonRenewable)
			saveErr = transactionApp.Save(fuelNonRenewableConsumptionEntry)
			if saveErr != nil {
				return saveErr
			}
			fileImport.Set("entires+", fuelNonRenewableConsumptionEntry)
			transactionApp.Save(fileImport)
		}

		if energyData.consumption.heat != nil {
			unitId, err := GetUnitId(transactionApp, energyData.consumption.heat.unit)
			if err != nil {
				return err
			}

			heatRenewableConsumptionEntry := core.NewRecord(entries)
			heatRenewableConsumptionEntry.Set("organization", organizationId)
			heatRenewableConsumptionEntry.Set("domain", heatDomain)
			heatRenewableConsumptionEntry.Set("type", renewableTypeId)
			heatRenewableConsumptionEntry.Set("aggregation", aggregationYearId)
			heatRenewableConsumptionEntry.Set("timestamp", energyData.startDate)
			heatRenewableConsumptionEntry.Set("unit", unitId)
			heatRenewableConsumptionEntry.Set("value", energyData.consumption.heat.renewable)
			saveErr := transactionApp.Save(heatRenewableConsumptionEntry)
			if saveErr != nil {
				return saveErr
			}

			heatNonRenewableConsumptionEntry := core.NewRecord(entries)
			heatNonRenewableConsumptionEntry.Set("organization", organizationId)
			heatNonRenewableConsumptionEntry.Set("domain", heatDomain)
			heatNonRenewableConsumptionEntry.Set("type", nonRenewableTypeId)
			heatNonRenewableConsumptionEntry.Set("aggregation", aggregationYearId)
			heatNonRenewableConsumptionEntry.Set("timestamp", energyData.startDate)
			heatNonRenewableConsumptionEntry.Set("unit", unitId)
			heatNonRenewableConsumptionEntry.Set("value", energyData.consumption.heat.nonRenewable)
			saveErr = transactionApp.Save(heatNonRenewableConsumptionEntry)
			if saveErr != nil {
				return saveErr
			}
			fileImport.Set("entires+", heatNonRenewableConsumptionEntry)
			transactionApp.Save(fileImport)
		}

		if energyData.emission.electricity != nil {
			unitId, err := GetUnitId(transactionApp, energyData.emission.electricity.unit)
			if err != nil {
				return err
			}

			electricityCO2EmissionEntry := core.NewRecord(entries)
			electricityCO2EmissionEntry.Set("organization", organizationId)
			electricityCO2EmissionEntry.Set("domain", electricityDomain)
			electricityCO2EmissionEntry.Set("type", co2EmissionTypeId)
			electricityCO2EmissionEntry.Set("aggregation", aggregationYearId)
			electricityCO2EmissionEntry.Set("timestamp", energyData.startDate)
			electricityCO2EmissionEntry.Set("unit", unitId)
			electricityCO2EmissionEntry.Set("value", energyData.emission.electricity.value)
			saveErr := transactionApp.Save(electricityCO2EmissionEntry)
			if saveErr != nil {
				return saveErr
			}
			fileImport.Set("entires+", electricityCO2EmissionEntry)
			transactionApp.Save(fileImport)
		}

		if energyData.emission.fuel != nil {
			unitId, err := GetUnitId(transactionApp, energyData.emission.fuel.unit)
			if err != nil {
				return err
			}

			fuelCO2EmissionEntry := core.NewRecord(entries)
			fuelCO2EmissionEntry.Set("organization", organizationId)
			fuelCO2EmissionEntry.Set("domain", fuelDomain)
			fuelCO2EmissionEntry.Set("type", co2EmissionTypeId)
			fuelCO2EmissionEntry.Set("aggregation", aggregationYearId)
			fuelCO2EmissionEntry.Set("timestamp", energyData.startDate)
			fuelCO2EmissionEntry.Set("unit", unitId)
			fuelCO2EmissionEntry.Set("value", energyData.emission.fuel.value)
			saveErr := transactionApp.Save(fuelCO2EmissionEntry)
			if saveErr != nil {
				return saveErr
			}
			fileImport.Set("entires+", fuelCO2EmissionEntry)
			transactionApp.Save(fileImport)
		}

		if energyData.emission.heat != nil {
			unitId, err := GetUnitId(transactionApp, energyData.emission.heat.unit)
			if err != nil {
				return err
			}

			heatCO2EmissionEntry := core.NewRecord(entries)
			heatCO2EmissionEntry.Set("organization", organizationId)
			heatCO2EmissionEntry.Set("domain", heatDomain)
			heatCO2EmissionEntry.Set("type", co2EmissionTypeId)
			heatCO2EmissionEntry.Set("aggregation", aggregationYearId)
			heatCO2EmissionEntry.Set("timestamp", energyData.startDate)
			heatCO2EmissionEntry.Set("unit", unitId)
			heatCO2EmissionEntry.Set("value", energyData.emission.heat.value)
			saveErr := transactionApp.Save(heatCO2EmissionEntry)
			if saveErr != nil {
				return saveErr
			}
			fileImport.Set("entires+", heatCO2EmissionEntry)
			transactionApp.Save(fileImport)
		}

		fileImport.Set("status", "Processed without errors")
		transactionApp.Save(fileImport)

		return nil
	})

	return createEnergyDataRecordsError
}

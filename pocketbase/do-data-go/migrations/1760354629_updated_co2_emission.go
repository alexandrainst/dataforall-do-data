package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1408555909")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT id\nFROM view_entries\nWHERE typeName = \"CO2 Emission\""
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_VLsI")

		// remove field
		collection.Fields.RemoveById("_clone_uj3d")

		// remove field
		collection.Fields.RemoveById("json3145888567")

		// remove field
		collection.Fields.RemoveById("_clone_aPn9")

		// remove field
		collection.Fields.RemoveById("_clone_hBKX")

		// remove field
		collection.Fields.RemoveById("_clone_xjVq")

		// remove field
		collection.Fields.RemoveById("_clone_nPPM")

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1408555909")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  value,\n  organizations.name as organization,\n  strftime('%Y', timestamp) as year,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName,\n  time_aggregations.name as timeAggregation\nFROM entries\nLEFT JOIN domains ON domains.id = entries.domain\nLEFT JOIN units ON units.id = entries.unit\nLEFT JOIN types ON types.id = entries.type\nLEFT JOIN time_aggregations ON time_aggregations.id = entries.aggregation\nLEFT JOIN organizations ON organizations.id = entries.organization\nWHERE types.name = \"CO2 Emission\""
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_VLsI",
			"max": null,
			"min": null,
			"name": "value",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_uj3d",
			"max": 400,
			"min": 0,
			"name": "organization",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "json3145888567",
			"maxSize": 1,
			"name": "year",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "json"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_aPn9",
			"max": 0,
			"min": 0,
			"name": "domainName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_hBKX",
			"max": 0,
			"min": 0,
			"name": "unitName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_xjVq",
			"max": 0,
			"min": 0,
			"name": "typeName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_nPPM",
			"max": 0,
			"min": 0,
			"name": "timeAggregation",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}

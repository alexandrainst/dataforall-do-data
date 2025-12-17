package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3582517491")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  value,\n  organizations.name as organization,\n  strftime('%Y', timestamp) as year,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName,\n  time_aggregations.name as timeAggregation\nFROM entries\nLEFT JOIN domains ON domains.id = entries.domain\nLEFT JOIN units ON units.id = entries.unit\nLEFT JOIN types ON types.id = entries.type\nLEFT JOIN time_aggregations ON time_aggregations.id = entries.aggregation\nLEFT JOIN organizations ON organizations.id = entries.organization"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_auNY")

		// remove field
		collection.Fields.RemoveById("_clone_7bsD")

		// remove field
		collection.Fields.RemoveById("_clone_yEia")

		// remove field
		collection.Fields.RemoveById("_clone_oBhp")

		// remove field
		collection.Fields.RemoveById("_clone_M5fQ")

		// remove field
		collection.Fields.RemoveById("_clone_ZsbA")

		// remove field
		collection.Fields.RemoveById("_clone_3i6w")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_qyS7",
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
			"id": "_clone_nGpZ",
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
			"id": "_clone_cy5p",
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
			"id": "_clone_ujdv",
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
			"id": "_clone_QWnN",
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
			"id": "_clone_09NY",
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
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3582517491")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  value,\n  organizations.name as organization,\n  timestamp,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName,\n  time_aggregations.name as timeAggregation\nFROM entries\nLEFT JOIN domains ON domains.id = entries.domain\nLEFT JOIN units ON units.id = entries.unit\nLEFT JOIN types ON types.id = entries.type\nLEFT JOIN time_aggregations ON time_aggregations.id = entries.aggregation\nLEFT JOIN organizations ON organizations.id = entries.organization"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_auNY",
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
			"id": "_clone_7bsD",
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
			"id": "_clone_yEia",
			"max": "",
			"min": "",
			"name": "timestamp",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "date"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_oBhp",
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
			"id": "_clone_M5fQ",
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
			"id": "_clone_ZsbA",
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
			"id": "_clone_3i6w",
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

		// remove field
		collection.Fields.RemoveById("_clone_qyS7")

		// remove field
		collection.Fields.RemoveById("_clone_nGpZ")

		// remove field
		collection.Fields.RemoveById("json3145888567")

		// remove field
		collection.Fields.RemoveById("_clone_cy5p")

		// remove field
		collection.Fields.RemoveById("_clone_ujdv")

		// remove field
		collection.Fields.RemoveById("_clone_QWnN")

		// remove field
		collection.Fields.RemoveById("_clone_09NY")

		return app.Save(collection)
	})
}

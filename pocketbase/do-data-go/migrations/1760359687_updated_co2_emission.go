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
			"viewQuery": "SELECT id, value, organizationName as organization, year, domainName, unitName, typeName, timeAggregation\nFROM view_entries\nWHERE typeName = \"CO2 Emission\""
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_YVjE")

		// remove field
		collection.Fields.RemoveById("_clone_Sfgp")

		// remove field
		collection.Fields.RemoveById("_clone_2syL")

		// remove field
		collection.Fields.RemoveById("_clone_zxVV")

		// remove field
		collection.Fields.RemoveById("_clone_X6DL")

		// remove field
		collection.Fields.RemoveById("_clone_qWla")

		// remove field
		collection.Fields.RemoveById("_clone_kabd")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_Pa7X",
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
			"id": "_clone_TOZu",
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
			"id": "_clone_pWpn",
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
			"id": "_clone_OYzp",
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
			"id": "_clone_W3Dt",
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
			"id": "_clone_h5Zp",
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
			"id": "_clone_dVXH",
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
		collection, err := app.FindCollectionByNameOrId("pbc_1408555909")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT id, value, organization, year, domainName, unitName, typeName, timeAggregation\nFROM view_entries\nWHERE typeName = \"CO2 Emission\""
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_YVjE",
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
			"cascadeDelete": false,
			"collectionId": "pbc_2387082370",
			"hidden": false,
			"id": "_clone_Sfgp",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "_clone_2syL",
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
			"id": "_clone_zxVV",
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
			"id": "_clone_X6DL",
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
			"id": "_clone_qWla",
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
			"id": "_clone_kabd",
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
		collection.Fields.RemoveById("_clone_Pa7X")

		// remove field
		collection.Fields.RemoveById("_clone_TOZu")

		// remove field
		collection.Fields.RemoveById("_clone_pWpn")

		// remove field
		collection.Fields.RemoveById("_clone_OYzp")

		// remove field
		collection.Fields.RemoveById("_clone_W3Dt")

		// remove field
		collection.Fields.RemoveById("_clone_h5Zp")

		// remove field
		collection.Fields.RemoveById("_clone_dVXH")

		return app.Save(collection)
	})
}

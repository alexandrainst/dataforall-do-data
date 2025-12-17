package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_436022809")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  comparison_entries.id,\n  comparison_entries.category,\n  value,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName\nFROM comparison_entries\nLEFT JOIN domains ON domains.id = comparison_entries.domain\nLEFT JOIN units ON units.id = comparison_entries.unit\nLEFT JOIN types ON types.id = comparison_entries.type"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_6o3y")

		// remove field
		collection.Fields.RemoveById("_clone_oJfy")

		// remove field
		collection.Fields.RemoveById("_clone_ZKME")

		// remove field
		collection.Fields.RemoveById("_clone_qbtS")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1478425883",
			"hidden": false,
			"id": "_clone_Sx26",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "category",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "_clone_36yG",
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
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_gP8K",
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_0e3b",
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
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_KgRR",
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

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_436022809")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  comparison_entries.id,\n  value,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName\nFROM comparison_entries\nLEFT JOIN domains ON domains.id = comparison_entries.domain\nLEFT JOIN units ON units.id = comparison_entries.unit\nLEFT JOIN types ON types.id = comparison_entries.type"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_6o3y",
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
			"id": "_clone_oJfy",
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
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_ZKME",
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_qbtS",
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

		// remove field
		collection.Fields.RemoveById("_clone_Sx26")

		// remove field
		collection.Fields.RemoveById("_clone_36yG")

		// remove field
		collection.Fields.RemoveById("_clone_gP8K")

		// remove field
		collection.Fields.RemoveById("_clone_0e3b")

		// remove field
		collection.Fields.RemoveById("_clone_KgRR")

		return app.Save(collection)
	})
}

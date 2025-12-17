package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2556637407")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_Jm3kYugNr0` + "`" + ` ON ` + "`" + `comparison_entries` + "`" + ` (\n  ` + "`" + `domain` + "`" + `,\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `year` + "`" + `\n)"
			],
			"listRule": "",
			"name": "comparison_entries",
			"viewRule": ""
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "number494360628",
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
			"collectionId": "pbc_2457221083",
			"hidden": false,
			"id": "relation2812878347",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "domain",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2195891175",
			"hidden": false,
			"id": "relation105650625",
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_24572210832",
			"hidden": false,
			"id": "relation2363381545",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "type",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_586599074",
			"hidden": false,
			"id": "relation3703245907",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "unit",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"hidden": false,
			"id": "number3145888567",
			"max": null,
			"min": null,
			"name": "year",
			"onlyInt": false,
			"presentable": false,
			"required": true,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2556637407")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [],
			"listRule": null,
			"name": "comparison_profiles",
			"viewRule": null
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("number494360628")

		// remove field
		collection.Fields.RemoveById("relation2812878347")

		// remove field
		collection.Fields.RemoveById("relation105650625")

		// remove field
		collection.Fields.RemoveById("relation2363381545")

		// remove field
		collection.Fields.RemoveById("relation3703245907")

		// remove field
		collection.Fields.RemoveById("number3145888567")

		return app.Save(collection)
	})
}

package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_500998982")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation3914652515")

		// remove field
		collection.Fields.RemoveById("relation2812878347")

		// remove field
		collection.Fields.RemoveById("relation2363381545")

		// remove field
		collection.Fields.RemoveById("relation3962411722")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1121724375",
			"hidden": false,
			"id": "relation1188605132",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "rule",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_500998982")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2195891175",
			"hidden": false,
			"id": "relation3914652515",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization_category",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
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
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2665740719",
			"hidden": false,
			"id": "relation3962411722",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "situation",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation1188605132")

		return app.Save(collection)
	})
}

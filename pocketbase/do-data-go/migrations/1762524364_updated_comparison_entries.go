package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2556637407")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1478425883",
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

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2556637407")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation105650625")

		return app.Save(collection)
	})
}

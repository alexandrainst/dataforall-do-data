package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1755329703")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  recommendation_templates.title,\n  recommendation_statistics.rule,\n  recommendation_statistics.organization,\n  recommendation_templates.text\nFROM recommendation_templates \nJOIN recommendation_statistics ON recommendation_templates.rule = recommendation_statistics.rule"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_i73j")

		// remove field
		collection.Fields.RemoveById("_clone_M8tK")

		// remove field
		collection.Fields.RemoveById("_clone_FfOX")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_SZp2",
			"max": 0,
			"min": 0,
			"name": "title",
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
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1121724375",
			"hidden": false,
			"id": "_clone_Z37H",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "rule",
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
			"collectionId": "pbc_2387082370",
			"hidden": false,
			"id": "_clone_NBzT",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "organization",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"convertURLs": false,
			"hidden": false,
			"id": "_clone_nFgB",
			"maxSize": 0,
			"name": "text",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "editor"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1755329703")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  recommendation_templates.title,\n  recommendation_statistics.rule,\n  recommendation_templates.text\nFROM recommendation_templates \nJOIN recommendation_statistics ON recommendation_templates.rule = recommendation_statistics.rule"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_i73j",
			"max": 0,
			"min": 0,
			"name": "title",
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
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_1121724375",
			"hidden": false,
			"id": "_clone_M8tK",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "rule",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"convertURLs": false,
			"hidden": false,
			"id": "_clone_FfOX",
			"maxSize": 0,
			"name": "text",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "editor"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_SZp2")

		// remove field
		collection.Fields.RemoveById("_clone_Z37H")

		// remove field
		collection.Fields.RemoveById("_clone_NBzT")

		// remove field
		collection.Fields.RemoveById("_clone_nFgB")

		return app.Save(collection)
	})
}

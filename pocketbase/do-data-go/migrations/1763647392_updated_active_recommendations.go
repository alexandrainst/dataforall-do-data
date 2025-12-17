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
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  recommendation_templates.title,\n  recommendation_statistics.rule,\n  recommendation_statistics.organization,\n  recommendation_templates.text,\n  rules.type\nFROM recommendation_templates \nJOIN recommendation_statistics ON recommendation_templates.rule = recommendation_statistics.rule\nJOIN rules ON rules.id = recommendation_statistics.rule"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_n96g")

		// remove field
		collection.Fields.RemoveById("_clone_e45a")

		// remove field
		collection.Fields.RemoveById("_clone_FxPW")

		// remove field
		collection.Fields.RemoveById("_clone_KrBh")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_ppPV",
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
			"id": "_clone_GNkN",
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
			"id": "_clone_BFPb",
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
			"id": "_clone_LzBG",
			"maxSize": 0,
			"name": "text",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "editor"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_24572210832",
			"hidden": false,
			"id": "_clone_znaT",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "type",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
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
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  recommendation_templates.title,\n  recommendation_statistics.rule,\n  recommendation_statistics.organization,\n  recommendation_templates.text\nFROM recommendation_templates \nJOIN recommendation_statistics ON recommendation_templates.rule = recommendation_statistics.rule\nJOIN rules ON rules.id = recommendation_statistics.rule"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_n96g",
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
			"id": "_clone_e45a",
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
			"id": "_clone_FxPW",
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
			"id": "_clone_KrBh",
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
		collection.Fields.RemoveById("_clone_ppPV")

		// remove field
		collection.Fields.RemoveById("_clone_GNkN")

		// remove field
		collection.Fields.RemoveById("_clone_BFPb")

		// remove field
		collection.Fields.RemoveById("_clone_LzBG")

		// remove field
		collection.Fields.RemoveById("_clone_znaT")

		return app.Save(collection)
	})
}

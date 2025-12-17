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
			"listRule": "  @request.auth.organization = organization"
		}`), &collection); err != nil {
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

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_cwgi",
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
			"id": "_clone_ipjp",
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
			"id": "_clone_pJ7n",
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
			"id": "_clone_2Cfw",
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
			"listRule": ""
		}`), &collection); err != nil {
			return err
		}

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

		// remove field
		collection.Fields.RemoveById("_clone_cwgi")

		// remove field
		collection.Fields.RemoveById("_clone_ipjp")

		// remove field
		collection.Fields.RemoveById("_clone_pJ7n")

		// remove field
		collection.Fields.RemoveById("_clone_2Cfw")

		return app.Save(collection)
	})
}

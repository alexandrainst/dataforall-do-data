package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_795917085")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@request.auth.organization = organization || @request.auth.role.type = \"MUNICIPALITY\""
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_9iu2")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2387082370",
			"hidden": false,
			"id": "_clone_jGGX",
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

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_795917085")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@request.auth.organization = organization || @request.auth.role.type = \"MUNICIAPLITY\""
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2387082370",
			"hidden": false,
			"id": "_clone_9iu2",
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

		// remove field
		collection.Fields.RemoveById("_clone_jGGX")

		return app.Save(collection)
	})
}

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
			"viewQuery": "SELECT id, value\nFROM view_entries\nWHERE typeName = \"CO2 Emission\""
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_2xe0",
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

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1408555909")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT id\nFROM view_entries\nWHERE typeName = \"CO2 Emission\""
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_2xe0")

		return app.Save(collection)
	})
}

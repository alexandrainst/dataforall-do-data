package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3165375535")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.organization = organization",
			"deleteRule": "@request.auth.organization = organization",
			"listRule": "@request.auth.organization = organization",
			"updateRule": "@request.auth.organization = organization",
			"viewRule": "@request.auth.organization = organization"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3165375535")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": null,
			"deleteRule": null,
			"listRule": "",
			"updateRule": null,
			"viewRule": ""
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}

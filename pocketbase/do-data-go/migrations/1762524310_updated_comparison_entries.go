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
			"indexes": []
		}`), &collection); err != nil {
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
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_Jm3kYugNr0` + "`" + ` ON ` + "`" + `comparison_entries` + "`" + ` (\n  ` + "`" + `domain` + "`" + `,\n  ` + "`" + `category` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `year` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}

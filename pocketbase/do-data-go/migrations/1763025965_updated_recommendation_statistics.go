package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2622782114")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_RFAdgKgFUL` + "`" + ` ON ` + "`" + `recommendation_statistics` + "`" + ` (\n  ` + "`" + `domain` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `situation` + "`" + `,\n  ` + "`" + `rule` + "`" + `,\n  ` + "`" + `organization` + "`" + `,\n  ` + "`" + `year` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2622782114")
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
	})
}

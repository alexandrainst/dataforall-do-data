package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		record := core.NewRecord(superusers)

		// note: the values can be eventually loaded via os.Getenv(key)
		// or from a special local config file
		record.Set("email", "admin@test.dk")
		record.Set("password", "admin1234")
		app.Save(record)

		return nil
	}, func(app core.App) error {
		record, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, "admin@test.com")
		if record == nil {
			return nil // probably already deleted
		}

		return app.Delete(record)
	})
}

package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
)

func init() {
	migrations.Register(func(app core.App) error {
		roles, err := app.FindCollectionByNameOrId("roles")
		if err != nil {
			return err
		}

		// Create municipality role
		municipality := core.NewRecord(roles)
		municipality.Set("type", "MUNICIPALITY")
		err = app.Save(municipality)
		if err != nil {
			return err
		}

		// Create regular user role
		user := core.NewRecord(roles)
		user.Set("type", "ORGANIZATION")
		err = app.Save(user)
		if err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		//	Remove municipality and user role
		roles, err := app.FindAllRecords("roles",
			dbx.In("type", "MUNICIPALITY", "ORGANIZATION"),
		)

		if err != nil {
			return err
		}

		for _, role := range roles {
			err := app.Delete(role)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

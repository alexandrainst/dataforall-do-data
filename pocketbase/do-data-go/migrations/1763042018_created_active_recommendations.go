package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text3208210256",
					"max": 0,
					"min": 0,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_dGHX",
					"max": 0,
					"min": 0,
					"name": "title",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_1121724375",
					"hidden": false,
					"id": "_clone_Ndus",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "rule",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				}
			],
			"id": "pbc_1755329703",
			"indexes": [],
			"listRule": "",
			"name": "active_recommendations",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  recommendation_templates.title,\n  recommendation_statistics.rule\nFROM recommendation_templates \nJOIN recommendation_statistics ON recommendation_templates.rule = recommendation_statistics.rule",
			"viewRule": ""
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_1755329703")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}

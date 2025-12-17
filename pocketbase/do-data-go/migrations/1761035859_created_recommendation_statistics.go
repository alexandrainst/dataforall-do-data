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
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_2195891175",
					"hidden": false,
					"id": "relation3914652515",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "organization_category",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_2457221083",
					"hidden": false,
					"id": "relation2812878347",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "domain",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_24572210832",
					"hidden": false,
					"id": "relation2363381545",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "type",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_2665740719",
					"hidden": false,
					"id": "relation3962411722",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "situation",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_1121724375",
					"hidden": false,
					"id": "relation1188605132",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "rule",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_500998982",
					"hidden": false,
					"id": "relation1214092105",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "recommendation_template",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_2387082370",
					"hidden": false,
					"id": "relation3253625724",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "organization",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"hidden": false,
					"id": "number3145888567",
					"max": null,
					"min": 1,
					"name": "year",
					"onlyInt": true,
					"presentable": false,
					"required": true,
					"system": false,
					"type": "number"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_2622782114",
			"indexes": [],
			"listRule": null,
			"name": "recommendation_statistics",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2622782114")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}

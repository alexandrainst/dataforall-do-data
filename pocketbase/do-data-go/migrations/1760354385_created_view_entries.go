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
					"hidden": false,
					"id": "_clone_tYDK",
					"max": null,
					"min": null,
					"name": "value",
					"onlyInt": false,
					"presentable": false,
					"required": false,
					"system": false,
					"type": "number"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_ymbV",
					"max": 400,
					"min": 0,
					"name": "organization",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "json3145888567",
					"maxSize": 1,
					"name": "year",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "json"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_OwWt",
					"max": 0,
					"min": 0,
					"name": "domainName",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_obOG",
					"max": 0,
					"min": 0,
					"name": "unitName",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_lJcD",
					"max": 0,
					"min": 0,
					"name": "typeName",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": false,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "_clone_IEwB",
					"max": 0,
					"min": 0,
					"name": "timeAggregation",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				}
			],
			"id": "pbc_3834198252",
			"indexes": [],
			"listRule": null,
			"name": "view_entries",
			"system": false,
			"type": "view",
			"updateRule": null,
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  value,\n  organizations.name as organization,\n  strftime('%Y', timestamp) as year,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName,\n  time_aggregations.name as timeAggregation\nFROM entries\nLEFT JOIN domains ON domains.id = entries.domain\nLEFT JOIN units ON units.id = entries.unit\nLEFT JOIN types ON types.id = entries.type\nLEFT JOIN time_aggregations ON time_aggregations.id = entries.aggregation\nLEFT JOIN organizations ON organizations.id = entries.organization",
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3834198252")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}

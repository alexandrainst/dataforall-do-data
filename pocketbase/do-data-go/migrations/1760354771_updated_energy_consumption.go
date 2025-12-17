package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3582517491")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  id,\n  value,\n  organization,\n  year,\n  domainName,\n  unitName,\n  typeName,\n  timeAggregation\nFROM view_entries\nWHERE typeName in (\"Energy Consumption Renewable\", \"Energy Consumption Non-renewable\")"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_jiSu")

		// remove field
		collection.Fields.RemoveById("_clone_ZBFz")

		// remove field
		collection.Fields.RemoveById("json3145888567")

		// remove field
		collection.Fields.RemoveById("_clone_uvRA")

		// remove field
		collection.Fields.RemoveById("_clone_uz4D")

		// remove field
		collection.Fields.RemoveById("_clone_qp9b")

		// remove field
		collection.Fields.RemoveById("_clone_j9kk")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_9EyB",
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

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_lOzJ",
			"max": 400,
			"min": 0,
			"name": "organization",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "_clone_zBxw",
			"maxSize": 1,
			"name": "year",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "json"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_8Nz9",
			"max": 0,
			"min": 0,
			"name": "domainName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_MG1j",
			"max": 0,
			"min": 0,
			"name": "unitName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_U4kY",
			"max": 0,
			"min": 0,
			"name": "typeName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_DnFD",
			"max": 0,
			"min": 0,
			"name": "timeAggregation",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3582517491")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"viewQuery": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  value,\n  organizations.name as organization,\n  strftime('%Y', timestamp) as year,\n  domains.name as domainName,\n  units.name as unitName,\n  types.name as typeName,\n  time_aggregations.name as timeAggregation\nFROM entries\nLEFT JOIN domains ON domains.id = entries.domain\nLEFT JOIN units ON units.id = entries.unit\nLEFT JOIN types ON types.id = entries.type\nLEFT JOIN time_aggregations ON time_aggregations.id = entries.aggregation\nLEFT JOIN organizations ON organizations.id = entries.organization\nWHERE types.name in (\"Energy Consumption Renewable\", \"Energy Consumption Non-renewable\")"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"hidden": false,
			"id": "_clone_jiSu",
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

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_ZBFz",
			"max": 400,
			"min": 0,
			"name": "organization",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"hidden": false,
			"id": "json3145888567",
			"maxSize": 1,
			"name": "year",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "json"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_uvRA",
			"max": 0,
			"min": 0,
			"name": "domainName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_uz4D",
			"max": 0,
			"min": 0,
			"name": "unitName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_qp9b",
			"max": 0,
			"min": 0,
			"name": "typeName",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "_clone_j9kk",
			"max": 0,
			"min": 0,
			"name": "timeAggregation",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("_clone_9EyB")

		// remove field
		collection.Fields.RemoveById("_clone_lOzJ")

		// remove field
		collection.Fields.RemoveById("_clone_zBxw")

		// remove field
		collection.Fields.RemoveById("_clone_8Nz9")

		// remove field
		collection.Fields.RemoveById("_clone_MG1j")

		// remove field
		collection.Fields.RemoveById("_clone_U4kY")

		// remove field
		collection.Fields.RemoveById("_clone_DnFD")

		return app.Save(collection)
	})
}

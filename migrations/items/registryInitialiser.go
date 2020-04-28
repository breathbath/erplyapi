package items

import "github.com/breathbath/erplyapi/migrations"

func AddItems(registry *migrations.Registry) {
	registry.RegisterMigration(Migration20200427092907{})
}

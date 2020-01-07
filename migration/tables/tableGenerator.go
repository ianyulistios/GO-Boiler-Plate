package tables

import (
	"GO-Boiler-Plate/databases"
	"fmt"
)

// ExecuteMigration of Table
func ExecuteMigration() {
	fmt.Print("------EXECUTE THE MIGRATION---------")
	TableSchema()
	fmt.Print("-----------THE MIGRATION IS SUCCESS-----------")
}

// TableSchema Represented Schema To Migrate The Model
func TableSchema() {
	databases.GetConnection().AutoMigrate(
		User{},
	)
}

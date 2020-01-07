package tables

import (
	"GO-Boiler-Plate/databases"
	"fmt"
)

// ExecuteMigration of Table
func ExecuteMigration() {
	fmt.Print("------EXECUTE THE MIGRATION---------\n")
	TableSchema()
	fmt.Print("-----------THE MIGRATION IS SUCCESS-----------\n")
}

// TableSchema Represented Schema To Migrate The Model
func TableSchema() {
	fmt.Print("Process To Migrate\n")
	databases.GetConnection().AutoMigrate(
		User{},
	)
	fmt.Print("End of Process\n")
}

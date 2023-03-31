package main

import (
	_ "github.com/go-sql-driver/mysql"

	"migrate/cmd"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// m := lib.Migrate()
	// m.Up()
	cmd.Execute()
	// fmt.Println("migration-cli")
}

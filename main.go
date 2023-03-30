package main

import (
	_ "github.com/go-sql-driver/mysql"

	"migrate/cmd"
	"migrate/lib"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cmd.Execute()
	m := lib.Migrate()
	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}

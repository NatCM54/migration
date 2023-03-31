package main

import (
	_ "github.com/go-sql-driver/mysql"

	"migrate/cmd"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cmd.Execute()
}

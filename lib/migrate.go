package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func Migrate() *migrate.Migrate {
	db := NewMySql()
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("cannot create mysql instant:", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot get current directory:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/database/migrations", dir),
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("cannot create migrate instant:", err)
	}
	return m
}

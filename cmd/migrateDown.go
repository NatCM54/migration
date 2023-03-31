package cmd

import (
	"fmt"
	"migrate/lib"

	"github.com/spf13/cobra"
)

var migrateDownCmd *cobra.Command

func init() {
	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "migrate from v2 to v1 ",
		Long:  `Command to downgrade database from v2 to v1`,
		Run: func(cmd *cobra.Command, args []string) {
			m := lib.Migrate()
			m.Down()

			fmt.Println("Running migrate down command")

			// db := NewMySql()

			// driver, err := mysql.WithInstance(db, &mysql.Config{})
			// if err != nil {
			// 	log.Fatal("cannot create mysql instant:", err)
			// }

			// dir, err := os.Getwd()
			// if err != nil {
			// 	log.Fatal("cannot get current directory:", err)
			// }

			// m, err := migrate.NewWithDatabaseInstance(
			// 	fmt.Sprintf("file://%s/database/migrations", dir),
			// 	"mysql",
			// 	driver,
			// )
			// if err != nil {
			// 	log.Fatal("cannot create migrate instant:", err)
			// }

			// if err = m.Down(); err != nil {
			// 	fmt.Printf("migrate down error: %v \n", err)
			// }
			// fmt.Printf("Migrate down done with success")
		},
	}
	migrateCmd.AddCommand(migrateDownCmd)
}

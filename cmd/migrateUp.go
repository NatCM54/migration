package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"migrate/lib"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate to v1 command",
		Long:  `Command to install version 1 of our application`,
		Run: func(cmd *cobra.Command, args []string) {
			m := lib.Migrate()
			m.Up()

			fmt.Println("Running migrate up command")

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

			// if err = m.Up(); err != nil {
			// 	fmt.Printf("migrate up error: %v \n", err)
			// }
			// fmt.Printf("Migrate up done with success")
		},
	}
	migrateCmd.AddCommand(migrateUpCmd)
}

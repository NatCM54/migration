package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"migrate/lib"
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"mgt"},
	Short:   "migrate cmd is used for database migration",
	Long:    `migrate cmd is used for database migration: migrate < up | down >`,
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]
		step, _ := strconv.Atoi(args[1])
		m := lib.Migrate()

		if action == "up" {
			m.Steps(step)
			fmt.Println("Running migrate up command")
		} else if action == "down" {
			m.Steps(step * (-1))
			fmt.Println("Running migrate down command")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

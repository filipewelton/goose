package internal

import (
	"fmt"
	"retail_workflow/internal/infrastructure/postgres"

	"github.com/spf13/cobra"
)

var Rollback = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := postgres.Connect(); err != nil {
			panic(err)
		}

		defer postgres.Disconnect()

		sql := readScriptContent("rollback")
		tx := postgres.Client.Exec(sql)

		if tx.Error != nil {
			panic(tx.Error)
		}

		fmt.Println("All migrations have been reversed")
	},
}

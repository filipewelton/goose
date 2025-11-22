package internal

import (
	"fmt"
	"retail_workflow/internal/infrastructure/postgres"

	"github.com/spf13/cobra"
)

var Push = &cobra.Command{
	Use:   "push",
	Short: "Push all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := postgres.Connect(); err != nil {
			panic(err)
		}

		defer postgres.Disconnect()

		sql := readScriptContent("push")
		tx := postgres.Client.Exec(sql)

		if tx.Error != nil {
			panic(tx.Error)
		}

		fmt.Println("The migrations were applied")
	},
}

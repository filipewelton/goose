package main

import (
	"retail_flow/internal/infrastructure/database"
	"retail_flow/internal/persistence/models"
	"retail_flow/internal/shared/lib"

	"github.com/spf13/cobra"
)

var postgres database.Postgres

func main() {
	lib.SetupEnvironmentVariables()
	root.AddCommand(up, down)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}

var root = &cobra.Command{
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var up = &cobra.Command{
	Use:   "up",
	Short: "Push migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := postgres.Connect(); err != nil {
			panic(err)
		}

		err := postgres.Client.AutoMigrate(&models.UserModel{})

		if err != nil {
			panic(err)
		} else if err = postgres.Disconnect(); err != nil {
			panic(err)
		}
	},
}

var down = &cobra.Command{
	Use:   "down",
	Short: "Revert migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := postgres.Connect(); err != nil {
			panic(err)
		}

		migrator := postgres.Client.Migrator()
		err := migrator.DropTable(&models.UserModel{})

		if err != nil {
			panic(err)
		} else if err = postgres.Disconnect(); err != nil {
			panic(err)
		}
	},
}

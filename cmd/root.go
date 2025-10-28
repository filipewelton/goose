package cmd

import (
	"goose/cmd/config"
	"goose/cmd/db"

	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Short: "Database management tool",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func setup() {
	config.Setup()
	db.Setup()
}

func Execute() {
	setup()

	Root.AddCommand(config.Config, db.DB)

	if err := Root.Execute(); err != nil {
		panic(err)
	}
}

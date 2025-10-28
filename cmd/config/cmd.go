package config

import (
	"goose/cmd/config/create"
	"goose/cmd/config/setup"

	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "Goose configuration",
	Long:  "Goose configuration handling",
}

func Setup() {
	create.Setup()
	Config.AddCommand(create.Create, setup.Setup)
}

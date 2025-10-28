package create

import (
	"encoding/json"
	"errors"
	"goose/internal/constants"
	"goose/internal/lib"
	"goose/internal/typings"
	"os"

	"github.com/spf13/cobra"
)

var overwriteFlag bool

var Create = &cobra.Command{
	Use:   "create",
	Short: "Create a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		data := parseData()

		var filename = os.Getenv("GOOSE_CONFIG_FILE")

		if !overwriteFlag {
			checkIfFileIsExists(filename)
		}

		err := os.WriteFile(filename, data, 0666)

		if err != nil {
			lib.ThrowError("Failed to create configuration file", err)
		}
	},
}

func parseData() []byte {
	data, err := json.MarshalIndent(typings.Config{
		Migrations: "./migrations",
		Schemas:    "./schemas",
		Database: typings.Database{
			Driver: constants.SQLITE,
			DSN:    "./goose.sqlite",
		},
	}, "", "  ")

	if err != nil {
		lib.ThrowError("Failed to parse configuration data", err)
	}

	return data
}

func checkIfFileIsExists(filename string) {
	stat, _ := os.Stat(filename)

	if stat != nil {
		lib.ThrowError(
			"Configuration file already exists: ",
			errors.New(filename+" already exists"),
		)
	}
}

func Setup() {
	Create.Flags().BoolVarP(
		&overwriteFlag,
		"overwrite",
		"o",
		false,
		"Overwrite file configuration",
	)
}

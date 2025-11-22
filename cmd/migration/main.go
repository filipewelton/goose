package main

import (
	"retail_workflow/cmd/migration/internal"
	"retail_workflow/internal/shared/environment"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Short: "Manage database migrations",
}

func main() {
	environment.LoadEnvironmentVariables()
	root.AddCommand(internal.Push, internal.Rollback)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}

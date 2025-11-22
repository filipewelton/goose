package main

import (
	"fmt"
	"os"
	"regexp"
	"retail_workflow/internal/infrastructure/postgres"
	"retail_workflow/internal/shared/environment"
)

var pattern = regexp.MustCompile(`(retail_workflow)/?.*`)

func main() {
	environment.LoadEnvironmentVariables()

	if err := postgres.Connect(); err != nil {
		panic(err)
	}

	defer postgres.Disconnect()

	path, _ := os.Getwd()
	replacement := "retail_workflow/cmd/seeder/script.sql"
	path = pattern.ReplaceAllString(path, replacement)
	raw, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	tx := postgres.Client.Exec(string(raw))

	if tx.Error != nil {
		panic(tx)
	}

	fmt.Println("Initial data entry completed successfully")
}

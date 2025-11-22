package internal

import (
	"fmt"
	"os"
	"regexp"
)

var pattern = regexp.MustCompile(`(retail_workflow)/?.*`)

func readScriptContent(scriptName string) string {
	replacement :=
		fmt.Sprintf("retail_workflow/cmd/migration/scripts/%s.sql", scriptName)
	path, _ := os.Getwd()
	path = pattern.ReplaceAllString(path, replacement)
	raw, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(raw)
}

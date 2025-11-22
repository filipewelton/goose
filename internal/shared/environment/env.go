package environment

import (
	"os"
	"regexp"
	"retail_workflow/internal/shared/errors"

	"github.com/joho/godotenv"
)

type Key string

const GOENV Key = "GOENV"
const HTTP_SERVER_ADDR Key = "HTTP_SERVER_ADDR"
const CACHE_ADDR Key = "CACHE_ADDR"
const CACHE_USERNAME Key = "CACHE_USERNAME"
const CACHE_PASSWORD Key = "CACHE_PASSWORD"
const LOGGING_ADDR Key = "LOGGING_ADDR"
const LOGGING_USERNAME Key = "LOGGING_USERNAME"
const LOGGING_PASSWORD Key = "LOGGING_PASSWORD"
const POSTGRES_DSN Key = "POSTGRES_DSN"
const API_GUARDIAN_SUBJECT Key = "API_GUARDIAN_SUBJECT"
const API_GUARDIAN_SECRET Key = "API_GUARDIAN_SECRET"
const USER_GUARDIAN_SECRET Key = "USER_GUARDIAN_SECRET"

var pattern = regexp.MustCompile(`(retail_workflow)/?.*`)

func GetEnv(k Key) string {
	return os.Getenv(string(k))
}

func LoadEnvironmentVariables() {
	goEnv := os.Getenv("GOENV")

	var err error

	switch goEnv {
	case "development", "test", "":
		err = godotenv.Overload(getPath(".env"))

		if err != nil {
			panic(err)
		}
	}

	checkEnvironmentVariables()
}

func getKeys() []string {
	return []string{
		string(HTTP_SERVER_ADDR),
		string(CACHE_ADDR),
		string(CACHE_USERNAME),
		string(CACHE_PASSWORD),
		string(LOGGING_ADDR),
		string(LOGGING_USERNAME),
		string(LOGGING_PASSWORD),
		string(POSTGRES_DSN),
		string(API_GUARDIAN_SUBJECT),
		string(API_GUARDIAN_SECRET),
		string(USER_GUARDIAN_SECRET),
	}
}

func UnloadEnvironmentVariables() {
	for _, v := range getKeys() {
		os.Setenv(v, "")
	}
}

func getPath(filename string) string {
	path, _ := os.Getwd()
	path = pattern.ReplaceAllString(path, "retail_workflow/"+filename)

	return path
}

func checkEnvironmentVariables() {
	for _, v := range getKeys() {
		if os.Getenv(v) == "" {
			panic(errors.Error500.ErrMissingEnvironmentVariable)
		}
	}
}

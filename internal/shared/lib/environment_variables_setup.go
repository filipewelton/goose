package lib

import (
	"os"
	"regexp"
	"retail_flow/internal/shared/exceptions"

	"github.com/joho/godotenv"
)

var variableNames = []string{
	"HTTP_SERVER_ADDR",
	"JWT_SECRET",
	"DATABASE_DSN",
	"REDIS_DSN",
	"REDIS_PASSWORD",
}

func SetupEnvironmentVariables() {
	goEnv := os.Getenv("GOENV")

	switch goEnv {
	case "production", "staging", "ci":
		break

	default:
		path, _ := os.Getwd()
		re := regexp.MustCompile(`(retail_workflow).*`)
		path = re.ReplaceAllString(path, "retail_workflow/.env")
		err := godotenv.Load(path)

		if err != nil {
			panic(err)
		}
	}

	checkVariables()
}

func checkVariables() {
	for _, name := range variableNames {
		if os.Getenv(name) == "" {
			panic(exceptions.InternalExceptions.ErrEnvironmentVariableIsMissing)
		}
	}
}

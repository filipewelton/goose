package lib

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func ReloadDotEnv() {
	path, _ := os.Getwd()
	re := regexp.MustCompile(`(retail_flow).*`)
	path = re.ReplaceAllString(path, "retail_flow/.env")
	err := godotenv.Load(path)

	if err != nil {
		panic(err)
	}
}

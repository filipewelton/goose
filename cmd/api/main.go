package main

import (
	"net/http"
	"os"
	httpDriver "retail_flow/internal/drivers/http"
	"retail_flow/internal/infrastructure/database"
	"retail_flow/internal/shared/lib"
	"time"
)

func main() {
	lib.SetupEnvironmentVariables()
	testDatabaseConnection()
	createHTTPServer()
}

func testDatabaseConnection() {
	var postgres database.Postgres

	if err := postgres.Connect(); err != nil {
		panic(err)
	}

	defer postgres.Disconnect()
}

func createHTTPServer() {
	var handler = httpDriver.SetupServer()
	var server = http.Server{
		Handler:      handler,
		Addr:         os.Getenv("HTTP_SERVER_ADDR"),
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

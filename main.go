package main

import (
	"backend-test/cmd"
	"backend-test/helpers"

	_ "github.com/lib/pq"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// run http
	cmd.ServeHTTP()
}

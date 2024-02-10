package main

import (
	"log"
	"os"
	"rinha_api/backend/config"
	"rinha_api/backend/httpd/route"
)

func init() {
	if err := config.ConnectInstance(); err != nil {
		log.Fatalf("err connect db %s", err.Error())
	}

	config.ExecMigration()
}

func main() {

	app := route.New().Routes()

	if err := app.Listen(os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Error listening on port %s: %s", os.Getenv("SERVER_PORT"), err)
	}

}

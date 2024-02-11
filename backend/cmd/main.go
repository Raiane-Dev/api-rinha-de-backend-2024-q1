package main

import (
	"log"
	"rinha_api/backend/config"
	"rinha_api/backend/httpd/route"
	"rinha_api/backend/util/logger"
)

const SERVER_PORT = ":80"

func init() {
	logger.Init()

	if err := config.ConnectInstance(); err != nil {
		log.Fatalf("err connect db %s", err.Error())
	}

	config.ExecMigration()
}

func main() {

	app := route.New().Routes()

	if err := app.Listen(SERVER_PORT); err != nil {
		log.Fatalf("Error listening on port %s: %s", SERVER_PORT, err)
	}

}

package route

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Router struct {
	Server  *fiber.App
	Clients fiber.Router
}

func New() *Router {
	log.Println("running routes")

	app := &Router{
		Server: fiber.New(),
	}
	app.Server.Use(compress.New())

	app.Clients = app.Server.Group("/clientes")

	return app
}

func (app *Router) Routes() *fiber.App {
	app.clientPublic()

	return app.Server
}

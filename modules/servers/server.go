package servers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	app *fiber.App
}

func StartServer() {
	server := &server{
		// Make a fiber faster
		app: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}

	v := server.app.Group("v1")
	modules := NewModules(v)

	// modules list
	modules.ProductsModule().Init()

	log.Println("server start on \":3000\"")
	server.app.Listen(":3000")
}

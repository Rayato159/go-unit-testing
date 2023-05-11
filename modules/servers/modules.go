package servers

import (
	"github.com/gofiber/fiber/v2"
)

type IModule interface {
	ProductsModule() IProductsModule
}

type module struct {
	router fiber.Router
}

func NewModules(r fiber.Router) IModule {
	return &module{
		router: r,
	}
}

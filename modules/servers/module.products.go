package servers

import (
	"github.com/Rayato159/go-unit-testing/modules/products/productsHandlers"
	"github.com/Rayato159/go-unit-testing/modules/products/productsRepositories"
	"github.com/Rayato159/go-unit-testing/modules/products/productsUsecases"
	"github.com/gofiber/fiber/v2"
)

type IProductsModule interface {
	Init()
	Repository() productsRepositories.IProductsRepository
	Usecase() productsUsecases.IProductsUsecase
	Handler() productsHandlers.IProductsHandler
}
type productModule struct {
	router     fiber.Router
	repsotiory productsRepositories.IProductsRepository
	usecase    productsUsecases.IProductsUsecase
	handler    productsHandlers.IProductsHandler
}

func (m *module) ProductsModule() IProductsModule {
	productsRepository := productsRepositories.NewProductsRepository()
	productsUsecase := productsUsecases.NewProductsUsecase(productsRepository)
	productsHandler := productsHandlers.NewProductHandler(productsUsecase)

	return &productModule{
		router:     m.router,
		repsotiory: productsRepository,
		usecase:    productsRepository,
		handler:    productsHandler,
	}
}
func (p *productModule) Init() {
	r := p.router.Group("products")
	r.Get("/:product_id", p.Handler().FindOneProduct)
}
func (p *productModule) Repository() productsRepositories.IProductsRepository {
	return p.repsotiory
}
func (p *productModule) Usecase() productsUsecases.IProductsUsecase {
	return p.usecase
}
func (p *productModule) Handler() productsHandlers.IProductsHandler {
	return p.handler
}

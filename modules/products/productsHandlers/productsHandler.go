package productsHandlers

import (
	"github.com/Rayato159/go-unit-testing/modules/products/productsUsecases"
	"github.com/gofiber/fiber/v2"
)

type IProductsHandler interface {
	FindOneProduct(c *fiber.Ctx) error
}

type productsHandler struct {
	productsUsecase productsUsecases.IProductsUsecase
}

func NewProductHandler(productsUsecase productsUsecases.IProductsUsecase) IProductsHandler {
	return &productsHandler{productsUsecase: productsUsecase}
}

func (h *productsHandler) FindOneProduct(c *fiber.Ctx) error {
	// Params request
	productId := c.Params("product_id")

	// Call a usecase
	product, err := h.productsUsecase.FindOneProduct(productId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			struct {
				Message string `json:"message"`
			}{
				Message: err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&product)
}

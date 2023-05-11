package productsRepositories

import (
	"fmt"

	"github.com/Rayato159/go-unit-testing/modules/products"
)

type IProductsRepository interface {
	FindOneProduct(productId string) (*products.Product, error)
}

type productsRepository struct{}

func NewProductsRepository() IProductsRepository {
	return &productsRepository{}
}

func (r *productsRepository) FindOneProduct(productId string) (*products.Product, error) {
	products := map[string]*products.Product{
		"P-000001": {
			Id:    "P-000001",
			Title: "Pudding",
			Price: 39,
		},
		"P-000002": {
			Id:    "P-000002",
			Title: "Coke",
			Price: 12,
		},
	}

	if products[productId] != nil {
		return products[productId], nil
	}
	return nil, fmt.Errorf("product_id: %s not found", productId)
}

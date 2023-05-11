package myTests

import (
	"testing"

	"github.com/Rayato159/go-unit-testing/modules/products"
	"github.com/Rayato159/go-unit-testing/modules/servers"
)

type testFindOneProduct struct {
	productId string
	isErr     bool
	expect    string
}

func TestFindOneProduct(t *testing.T) {
	tests := []testFindOneProduct{
		{
			productId: "P-000003",
			isErr:     true,
			expect:    "product_id: P-000003 not found",
		},
		{
			productId: "P-000001",
			isErr:     false,
			expect: CompressToJSON(&products.Product{
				Id:    "P-000001",
				Title: "Pudding",
				Price: 39,
			}),
		},
	}

	productModule := servers.NewModules(nil).ProductsModule()
	for _, test := range tests {
		if test.isErr {
			_, err := productModule.Usecase().FindOneProduct(test.productId)
			if err.Error() != test.expect {
				t.Errorf("expect: %v, got: %v", test.expect, err.Error())
			}
		} else {
			result, err := productModule.Usecase().FindOneProduct(test.productId)
			if err != nil {
				t.Errorf("expect: %v, got: %v", nil, err)
			}
			if CompressToJSON(&result) != test.expect {
				t.Errorf("expect: %v, got: %v", test.expect, CompressToJSON(&result))
			}
		}
	}
}

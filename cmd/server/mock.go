package server

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
)

type testCase struct {
	name               string
	mocksConfig        *mocksConfig
	body               interface{}
	expectedStatusCode int
}

type mocksConfig struct {
	error
}

type fakeService struct {
	err error
}

func newFakeService(err error) *fakeService {
	return &fakeService{
		err,
	}
}

func (f *fakeService) CreateProduct(domain.Product) error {
	return f.err
}

func getProduct() *domain.Product {
	return &domain.Product{Sku: "FAL-1000000", Name: "name", Size: "size", Brand: "brand", Price: 100, PrincipalImage: "https://image"}
}

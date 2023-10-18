package products

import "api-products-maribel-martinez/pkg/api/app/domain"

type fakeDbRepo struct {
	errCreate error
}

func newFakeDbRepo(errCreate error) *fakeDbRepo {
	return &fakeDbRepo{errCreate: errCreate}
}

func (db *fakeDbRepo) CreateProduct(domain.Product) error {
	return db.errCreate
}

func getProduct() *domain.Product {
	return &domain.Product{Sku: "FAL-10000000001", Name: "name", Size: "size", Brand: "brand", Price: 100, PrincipalImage: "https://image"}
}

type testCase struct {
	name          string
	mocksConfig   *mocksConfig
	body          *domain.Product
	expectedValue error
}

type mocksConfig struct {
	error
}

package products

import "api-products-maribel-martinez/pkg/api/app/domain"

type fakeDbRepo struct {
	errCreate error
	errSearh  error
	errGet    error
	product   *domain.Product
	products  []domain.Product
}

func newFakeDbRepo(errCreate, errSearch, errGet error, product *domain.Product, products []domain.Product) *fakeDbRepo {
	return &fakeDbRepo{
		errCreate: errCreate,
		errSearh:  errSearch,
		errGet:    errGet,
		product:   product,
		products:  products,
	}
}

func (db *fakeDbRepo) CreateProduct(domain.Product) error {
	return db.errCreate
}

func (db *fakeDbRepo) SearchProduct(sku string) (*domain.Product, error) {
	return db.product, db.errSearh
}
func (db *fakeDbRepo) GetProducts() ([]domain.Product, error) {
	return db.products, db.errGet
}
func getProduct() *domain.Product {
	return &domain.Product{Sku: "FAL-10000000001", Name: "name", Size: "size", Brand: "brand", Price: 100, PrincipalImage: "https://image"}
}

type testCaseCreate struct {
	name          string
	mocksConfig   *fakeDbRepo
	body          *domain.Product
	expectedValue error
}
type testCaseSearch struct {
	name          string
	mocksConfig   *fakeDbRepo
	sku           string
	expectedError error
	expectedValue *domain.Product
}
type testCaseGet struct {
	name          string
	mocksConfig   *fakeDbRepo
	sku           string
	expectedError error
	expectedValue []domain.Product
}

type mocksConfig struct {
	error
}

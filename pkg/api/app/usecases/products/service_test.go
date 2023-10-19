package products

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateProduct(t *testing.T) {
	testCases := []testCaseCreate{
		{
			"Ok",
			&fakeDbRepo{},
			getProduct(),
			nil,
		},
		{
			"error_bd",
			&fakeDbRepo{errCreate: errors.New("error")},
			getProduct(),
			errors.New("error creating product"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(y *testing.T) {
			dbRepo := newFakeDbRepo(tc.mocksConfig.errCreate, tc.mocksConfig.errSearh, tc.mocksConfig.errGet, tc.mocksConfig.product, tc.mocksConfig.products)
			service := NewService(dbRepo)
			err := service.CreateProduct(*tc.body)

			assert.Equal(t, tc.expectedValue, err)
		})
	}
}

func TestService_SearchProduct(t *testing.T) {
	testCases := []testCaseSearch{
		{
			"Ok",
			&fakeDbRepo{product: getProduct()},
			"sku-test",
			nil,
			getProduct(),
		},
		{
			"error_bd",
			&fakeDbRepo{errSearh: errors.New("error")},
			"sku-test",
			errors.New("error"),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(y *testing.T) {
			dbRepo := newFakeDbRepo(tc.mocksConfig.errCreate, tc.mocksConfig.errSearh, tc.mocksConfig.errGet, tc.mocksConfig.product, tc.mocksConfig.products)
			service := NewService(dbRepo)

			product, err := service.SearchProduct(tc.sku)

			assert.Equal(t, tc.expectedValue, product)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_GetProduct(t *testing.T) {
	testCases := []testCaseGet{
		{
			"Ok",
			&fakeDbRepo{products: []domain.Product{*getProduct()}},
			"sku-test",
			nil,
			[]domain.Product{*getProduct()},
		},
		{
			"error_bd",
			&fakeDbRepo{errGet: errors.New("error")},
			"sku-test",
			errors.New("error"),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(y *testing.T) {
			dbRepo := newFakeDbRepo(tc.mocksConfig.errCreate, tc.mocksConfig.errSearh, tc.mocksConfig.errGet, tc.mocksConfig.product, tc.mocksConfig.products)
			service := NewService(dbRepo)

			product, err := service.GetProducts()

			assert.Equal(t, tc.expectedValue, product)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

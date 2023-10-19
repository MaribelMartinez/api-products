package products

import "api-products-maribel-martinez/pkg/api/app/domain"

type DBRepo interface {
	CreateProduct(product domain.Product) error
	SearchProduct(sku string) (*domain.Product, error)
	GetProducts() ([]domain.Product, error)
}

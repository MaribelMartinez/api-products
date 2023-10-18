package server

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
)

type productService interface {
	CreateProduct(product domain.Product) error
}

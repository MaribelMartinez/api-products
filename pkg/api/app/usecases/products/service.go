package products

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"api-products-maribel-martinez/pkg/api/app/errors"
	"strings"
)

func NewService(dbRepo DBRepo) *Service {
	return &Service{dbRepo}
}

func (s *Service) CreateProduct(product domain.Product) error {
	product.Sku = strings.ToUpper(product.Sku)
	if err := s.DBRepo.CreateProduct(product); err != nil {
		return errors.ErrorFromMessage{Message: "error creating product"}
	}

	return nil
}

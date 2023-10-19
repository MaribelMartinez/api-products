package products

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"errors"
	"strings"
)

func NewService(dbRepo DBRepo) *Service {
	return &Service{dbRepo}
}

func (s *Service) CreateProduct(product domain.Product) error {
	product.Sku = strings.ToUpper(product.Sku)
	if err := s.DBRepo.CreateProduct(product); err != nil {
		return errors.New("error creating product")
	}
	return nil
}

func (s *Service) SearchProduct(sku string) (*domain.Product, error) {
	product, err := s.DBRepo.SearchProduct(sku)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) GetProducts() ([]domain.Product, error) {
	products, err := s.DBRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

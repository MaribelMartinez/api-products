package database

import "api-products-maribel-martinez/pkg/api/app/domain"

func getProduct() *domain.Product {
	return &domain.Product{Sku: "FAL-10000000001", Name: "name", Size: "size", Brand: "brand", Price: 100, PrincipalImage: "https://image"}
}

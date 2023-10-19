package database

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	errorsC "api-products-maribel-martinez/pkg/api/app/errors"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func NewRepository(db *gorm.DB) *DBConnection {
	return &DBConnection{connection: db}
}

func (db *DBConnection) CreateProduct(product domain.Product) error {
	return db.connection.Create(&product).Error
}

func (db *DBConnection) SearchProduct(sku string) (*domain.Product, error) {
	value, exists := db.connection.Get(sku)
	if !exists {
		return nil, errorsC.ErrorFromMessage{Message: "product don't exists", StatusCode: http.StatusNotFound}
	}
	product, ok := value.(domain.Product)
	if !ok {
		return nil, errors.New("internal error")
	}
	return &product, nil
}

func (db *DBConnection) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	result := db.connection.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

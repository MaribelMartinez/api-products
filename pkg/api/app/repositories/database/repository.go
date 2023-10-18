package database

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *DBConnection {
	return &DBConnection{connection: db}
}

func (db *DBConnection) CreateProduct(product domain.Product) error {
	return db.connection.Create(&product).Error
}

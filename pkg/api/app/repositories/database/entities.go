package database

import "gorm.io/gorm"

type DBConnection struct {
	connection *gorm.DB
}

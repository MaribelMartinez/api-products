package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestDBConnection_CreateProduct(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := mysql.New(mysql.Config{
		Conn:       mockDb,
		DriverName: "mysql",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	repository := NewRepository(db)

	err := repository.CreateProduct(*getProduct())

	assert.Nil(t, err)
}

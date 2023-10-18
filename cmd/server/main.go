package server

import (
	"api-products-maribel-martinez/pkg/api/app/repositories/database"
	"api-products-maribel-martinez/pkg/api/app/usecases/products"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("error initializing the application")
	}

}

func run() error {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("error open DB")
	}
	DBRepo := database.NewRepository(db)
	service := products.NewService(DBRepo)
	router := gin.Default()
	router.POST("/product", createProduct(service))

	if err := router.Run("localhost:8080"); err != nil {
		return errors.New("error to run localhost:8080")
	}
	return nil
}

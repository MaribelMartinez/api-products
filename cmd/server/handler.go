package server

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"api-products-maribel-martinez/pkg/api/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createProduct(service productService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body domain.Product
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		if err := helpers.ValidateProduct(body); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		if err := service.CreateProduct(body); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusCreated, "Product created")
		return
	}
}

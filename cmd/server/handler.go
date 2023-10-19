package server

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	errorsC "api-products-maribel-martinez/pkg/api/app/errors"
	"api-products-maribel-martinez/pkg/api/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createProduct(service productService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body domain.Product
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, errorsC.ErrorFromMessage{Message: err.Error(), StatusCode: http.StatusBadRequest})
			return
		}
		if err := helpers.ValidateProduct(body); err != nil {
			c.JSON(http.StatusBadRequest, errorsC.ErrorFromMessage{Message: err.Error(), StatusCode: http.StatusBadRequest})
			return
		}
		if err := service.CreateProduct(body); err != nil {
			c.JSON(http.StatusInternalServerError, errorsC.ErrorFromMessage{Message: err.Error(), StatusCode: http.StatusInternalServerError})
			return
		}
		c.Status(http.StatusCreated)
		return
	}
}

func searchProduct(service productService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sku := c.Query("sku")
		product, err := service.SearchProduct(sku)
		if err != nil {
			er := &errorsC.ErrorFromMessage{}
			if errors.As(err, &er) {
				c.JSON(er.StatusCode, er)
				return
			}
			c.JSON(http.StatusInternalServerError, errorsC.ErrorFromMessage{Message: err.Error(), StatusCode: http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, product)
		return
	}
}

func getProducts(service productService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sku := c.Query("sku")
		if sku == "" {
			c.JSON(http.StatusBadRequest, errorsC.ErrorFromMessage{Message: "SKU empty", StatusCode: http.StatusBadRequest})
			return
		}
		products, err := service.SearchProduct(sku)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorsC.ErrorFromMessage{Message: err.Error(), StatusCode: http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}
}

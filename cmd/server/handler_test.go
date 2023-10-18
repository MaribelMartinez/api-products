package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []testCase{
		{"status_created",
			&mocksConfig{},
			getProduct(),
			http.StatusCreated,
		},
		{"unmarshal_error",
			&mocksConfig{},
			gofight.D{"sku": 1, "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"SKU_error_min_length",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100.0, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"SKU_error_max_length",
			&mocksConfig{},
			gofight.D{"sku": "FAL-10000000000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"SKU_error_format_string",
			&mocksConfig{},
			gofight.D{"sku": "FA-10000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"SKU_error_format_number",
			&mocksConfig{},
			gofight.D{"sku": "FAL-00000e0", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"SKU_error_number",
			&mocksConfig{},
			gofight.D{"sku": "FAL-0000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"name_blank_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "   ", "brand": "Marca", "size": "M", "Price": 100, "prinicipal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"brand_blank_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "Zapatilla Mujer", "brand": " ", "size": "M", "Price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"size_blank_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "  ", "Price": 100, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"price_min_value_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 0, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"price_max_value_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 999999999, "principal_image": "https://images"},
			http.StatusBadRequest,
		},
		{"principal_image_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100, "principal_image": "image"},
			http.StatusBadRequest,
		},
		{"other_images_error",
			&mocksConfig{},
			gofight.D{"sku": "FAL-1000000", "name": "Zapatilla Mujer", "brand": "Marca", "size": "M", "Price": 100, "principal_image": "https://images", "other_images": []string{"https://", "https://images"}},
			http.StatusBadRequest,
		},
		{"internal_error",
			&mocksConfig{errors.New("error")},
			getProduct(),
			http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(y *testing.T) {
			router := gin.Default()
			service := newFakeService(tc.mocksConfig.error)
			router.POST("/product", createProduct(service))
			body, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/product", io.NopCloser(bytes.NewReader(body)))
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(y, tc.expectedStatusCode, rec.Code)
		})
	}
}

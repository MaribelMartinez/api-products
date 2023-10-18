package products

import (
	errorsC "api-products-maribel-martinez/pkg/api/app/errors"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateProduct(t *testing.T) {
	testCases := []testCase{
		{
			"Ok",
			&mocksConfig{nil},
			getProduct(),
			nil,
		},
		{
			"error_bd",
			&mocksConfig{errors.New("error")},
			getProduct(),
			errorsC.ErrorFromMessage{Message: "error creating product"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(y *testing.T) {
			dbRepo := newFakeDbRepo(tc.mocksConfig.error)
			service := NewService(dbRepo)

			err := service.CreateProduct(*tc.body)

			assert.Equal(t, tc.expectedValue, err)
		})
	}
}

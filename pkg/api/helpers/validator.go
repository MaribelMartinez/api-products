package helpers

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"api-products-maribel-martinez/pkg/api/app/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	url1 "net/url"
	"strings"
)

func validateProduct(product domain.Product) error {
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		return err
	}
	if err := validateBlankValuesProduct(product); err != nil {
		return err
	}
	if err := validateImagesUrls(product); err != nil {
		return err
	}
}

func validateBlank(value string) bool {
	return len(strings.TrimSpace(value)) == len(value)
}

func validateUrlFormat(url string) bool {
	u, err := url1.Parse(url)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateBlankValuesProduct(product domain.Product) error {
	if validateBlank(product.Name) {
		return errors.ErrorFromMessage{Message: "Name can't be blank"}
	}
	if validateBlank(product.Brand) {
		return errors.ErrorFromMessage{Message: "Brand can't be blank"}
	}
	if validateBlank(product.Size) {
		return errors.ErrorFromMessage{Message: "Size can't be blank"}

	}
	return nil
}

func validateImagesUrls(product domain.Product) error {
	if !validateUrlFormat(product.PrincipalImage) {
		return errors.ErrorFromMessage{Message: "the url of the principal image is not valid"}
	}
	for i, image := range product.OtherImages {
		if !validateUrlFormat(image) {
			return errors.ErrorFromMessage{Message: fmt.Sprintf("the url of other image in position %d is not valid", i)}
		}
	}
}

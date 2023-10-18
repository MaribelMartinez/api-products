package helpers

import (
	"api-products-maribel-martinez/pkg/api/app/domain"
	"api-products-maribel-martinez/pkg/api/app/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	url1 "net/url"
	"strconv"
	"strings"
)

func ValidateProduct(product domain.Product) error {
	validate := validator.New()
	var err error
	if err = validate.Struct(product); err != nil {
		return err
	}
	if err = validateSku(product.Sku); err != nil {
		return err
	}
	if err = validateBlankValuesProduct(product); err != nil {
		return err
	}
	if err = validateImagesUrls(product); err != nil {
		return err
	}
	return nil
}

func validateBlank(value string) bool {
	blank := len(strings.Trim(value, " "))
	return blank == 0
}

func validateUrlFormat(url string) bool {
	u, err := url1.Parse(url)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateSku(sku string) error {
	subStr := sku[0:4]
	subStrNum, err := strconv.Atoi(sku[4:])
	if err == nil {
		if subStr == domain.SubStrSku && subStrNum >= domain.MinSku && subStrNum <= domain.MaxSku {
			return nil
		}
	}

	return errors.ErrorFromMessage{Message: "SKU invalid"}
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
	return nil
}

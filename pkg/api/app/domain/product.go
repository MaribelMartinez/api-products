package domain

type Product struct {
	Sku            string   `json:"sku"`
	Name           string   `json:"name" validate:"required,max=50,min=3"`
	Brand          string   `json:"brand" validate:"required,max=50,min=3"`
	Size           string   `json:"size"`
	Price          float64  `json:"price" validate:"required"`
	PrincipalImage string   `json:"principal_image" validate:"required"`
	OtherImages    []string `json:"other_images"`
}

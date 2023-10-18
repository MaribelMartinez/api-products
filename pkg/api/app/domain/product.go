package domain

type Product struct {
	Sku            string   `json:"sku" validate:"required,max=12,min=11"`
	Name           string   `json:"name" validate:"required,max=50,min=3"`
	Brand          string   `json:"brand" validate:"required,max=50,min=3"`
	Size           string   `json:"size,omitempty"`
	Price          float64  `json:"price" validate:"required,gte=1,max=99999999"`
	PrincipalImage string   `json:"principal_image" validate:"required"`
	OtherImages    []string `json:"other_images,omitempty" gorm:"type:string[]"`
}

const (
	SubStrSku = "FAL-"
	MinSku    = 1000000
	MaxSku    = 99999999
)

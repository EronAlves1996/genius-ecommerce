package product

type Product struct {
	Id       uint32  `db:"id" json:"id"`
	Name     string  `validate:"required" json:"name" db:"name"`
	Price    float64 `validate:"required,gt=0" json:"price" db:"price"`
	ImageUrl string  `validate:"required" json:"image_url" db:"image_url"`
}

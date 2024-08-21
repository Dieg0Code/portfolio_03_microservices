package request

// CreateProductRequest struct
type CreateProductRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Category string `json:"category" validate:"required"`
	Price    int    `json:"price" validate:"required,min=1"`
	Stock    int    `json:"stock" validate:"required,min=1"`
}

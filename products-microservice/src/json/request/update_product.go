package request

// UpdateProductRequest struct
type UpdateProductRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=100"`
	Price int    `json:"price" validate:"required,min=1"`
	Stock int    `json:"stock" validate:"required,min=1"`
}

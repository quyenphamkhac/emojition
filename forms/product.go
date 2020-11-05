package forms

// CreateProductInput form
type CreateProductInput struct {
	Code  string `json:"code" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}

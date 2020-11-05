package forms

// SignUpInput ...
type SignUpInput struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

package forms

// SignUpInput ...
type SignUpInput struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

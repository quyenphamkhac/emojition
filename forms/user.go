package forms

// UserSignUp form
type UserSignUp struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Dob       string `json:"dob" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

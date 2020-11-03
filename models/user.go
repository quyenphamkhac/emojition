package models

// User model
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	Active    bool   `json:"active,omitempty"`
}

// SignUp function
// func (u User) SignUp(payload forms.UserSignUp) (*User, error) {

// }

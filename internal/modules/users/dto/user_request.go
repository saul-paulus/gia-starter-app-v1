package dto

// CreateUser is the request payload for creating a new user.
// @Description Request body for POST /users
type CreateUser struct {
	// Username must be between 3 and 100 characters
	// example: john_doe
	Username string `json:"username" binding:"required,min=3,max=100" example:"john_doe"`
	// Email must be a valid address and will be used as a unique identifier
	// example: john@example.com
	Email string `json:"email" binding:"required,email,max=254" example:"john@example.com"`
	// RoleID is the role identifier (1=admin, 2=user)
	// example: 2
	RoleID int `json:"role_id" binding:"required,min=1,max=2" example:"2"`
	// Password must be at least 8 characters and will be hashed with bcrypt
	// example: secretpassword
	Password string `json:"password" binding:"required,min=8,max=255" example:"secretpassword"`
}

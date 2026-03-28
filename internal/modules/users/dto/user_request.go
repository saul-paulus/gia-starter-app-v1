package dto

type CreateUser struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Email    string `json:"email" binding:"required,email,max=254"`
	RoleID   int    `json:"role_id" binding:"required,min=1,max=2"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

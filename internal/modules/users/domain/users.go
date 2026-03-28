package domain

import "time"

type Users struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"` // tidak dikirim ke response
	RoleID   int    `gorm:"column:id_role" json:"role_id"`

	IsActive bool `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

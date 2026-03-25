package seeder

import (
	"fmt"
	"gia-starter-app-V1/internal/modules/users/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	var user domain.Users

	// Check if default user already exists
	email := "saulpaulus17@gmail.com"
	err := db.Where("email = ?", email).First(&user).Error
	if err == nil {
		fmt.Printf("User with email %s already exists, skipping...\n", email)
		return nil
	}

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	// Hash password
	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Insert default user data
	newUser := domain.Users{
		Username: "saul paulus",
		Email:    email,
		Password: string(hashedPassword),
		RoleID:   1,
		IsActive: true,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return fmt.Errorf("failed to create default user: %w", err)
	}

	fmt.Println("✅ Default user created")
	return nil
}


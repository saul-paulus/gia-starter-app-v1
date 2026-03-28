package services

import (
	"errors"
	"gia-starter-app-V1/internal/modules/users/domain"
	"gia-starter-app-V1/internal/modules/users/dto"
	"gia-starter-app-V1/internal/modules/users/repositories"
	appErr "gia-starter-app-V1/internal/shared/errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersService interface {
	CreateUser(req dto.CreateUser) error
}

type usersService struct {
	repo repositories.UsersRepository
}

func NewUsersService(repo repositories.UsersRepository) UsersService {
	return &usersService{repo: repo}
}

// CreateUser membuat user baru dengan validasi email unik dan password yang di-hash
func (u *usersService) CreateUser(req dto.CreateUser) error {
	// Cek apakah email sudah terdaftar
	existing, err := u.repo.FindByEmailUser(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Error selain "record not found" dianggap error internal
		return appErr.ErrInternal
	}

	if existing != nil && existing.ID != 0 {
		return appErr.NewAppError(400, "EMAIL_EXISTS", "Email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return appErr.ErrInternal
	}

	user := domain.Users{
		Username: req.Username,
		Email:    req.Email,
		RoleID:   req.RoleID,
		Password: string(hashedPassword),
	}

	if err := u.repo.CreateUser(&user); err != nil {
		return appErr.ErrInternal
	}

	return nil
}

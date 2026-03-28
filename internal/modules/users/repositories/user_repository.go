package repositories

import (
	"gia-starter-app-V1/internal/modules/users/domain"

	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(user *domain.Users) error
	FindByEmailUser(email string) (*domain.Users, error)
}

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{db: db}
}

func (r *UsersRepositoryImpl) CreateUser(user *domain.Users) error {
	return r.db.Create(user).Error
}

func (r *UsersRepositoryImpl) FindByEmailUser(email string) (*domain.Users, error) {
	var user domain.Users

	err := r.db.Where("email =?", email).First(&user).Error
	return &user, err
}

package repository

import (
	"context"
	"gia-starter-app-V1/internal/modules/user/domain"
	"gia-starter-app-V1/internal/modules/user/domain/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.Users) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.Users, error) {
	var users []entity.Users
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.Users, error) {
	var user entity.Users
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.Users) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Users{}, id).Error
}

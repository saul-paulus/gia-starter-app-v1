package repository

import (
	"context"
	"gia-starter-app-V1/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.Users) error
	FindAll(ctx context.Context) ([]entity.Users, error)
	FindByID(ctx context.Context, id int) (*entity.Users, error)
	Update(ctx context.Context, user *entity.Users) error
	Delete(ctx context.Context, id int) error
}

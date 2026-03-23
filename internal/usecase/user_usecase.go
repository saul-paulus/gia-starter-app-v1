package usecase

import (
	"context"
	"gia-starter-app-V1/internal/domain/entity"
	"gia-starter-app-V1/internal/domain/repository"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.Users) error
	GetAllUsers(ctx context.Context) ([]entity.Users, error)
	GetUserByID(ctx context.Context, id int) (*entity.Users, error)
	UpdateUser(ctx context.Context, user *entity.Users) error
	DeleteUser(ctx context.Context, id int) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *entity.Users) error {
	return u.userRepo.Create(ctx, user)
}

func (u *userUseCase) GetAllUsers(ctx context.Context) ([]entity.Users, error) {
	return u.userRepo.FindAll(ctx)
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int) (*entity.Users, error) {
	return u.userRepo.FindByID(ctx, id)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *entity.Users) error {
	return u.userRepo.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int) error {
	return u.userRepo.Delete(ctx, id)
}

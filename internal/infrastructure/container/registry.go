package container

import (
	"gia-starter-app-V1/internal/infrastructure/persistence/postgres"
	"gia-starter-app-V1/internal/usecase"

	"gorm.io/gorm"
)

type Registry struct {
	UserUseCase usecase.UserUseCase
}

func NewRegistry(db *gorm.DB) *Registry {
	userRepo := postgres.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)

	return &Registry{
		UserUseCase: userUC,
	}
}

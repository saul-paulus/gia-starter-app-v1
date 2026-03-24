package container

import (
	"gia-starter-app-V1/internal/modules/user/infrastructure/persistence/postgres"
	"gia-starter-app-V1/internal/modules/user/usecase"

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

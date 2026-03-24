package user

import (
	"gorm.io/gorm"
	"gia-starter-app-V1/internal/modules/user/interface/http"
	"gia-starter-app-V1/internal/modules/user/interface/repository"
	"gia-starter-app-V1/internal/modules/user/usecase"
)

type Module struct {
	Handler *http.UserHandler
}

func Init(db *gorm.DB) *Module {
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	handler := http.NewUserHandler(uc)

	return &Module{
		Handler: handler,
	}
}

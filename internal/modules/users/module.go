package users

import (
	"gia-starter-app-V1/internal/modules/users/http"
	"gia-starter-app-V1/internal/modules/users/repositories"
	"gia-starter-app-V1/internal/modules/users/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	handler *http.UsersHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repositories.NewUsersRepository(db)
	svc := services.NewUsersService(repo)
	handler := http.NewUsersHandler(svc)

	return &Module{handler: handler}
}

func (m *Module) Register(r *gin.RouterGroup) {
	group := r.Group("/users")

	group.GET("", m.handler.Index)
	group.POST("", m.handler.CreateUserHandler)
}

package auth

import (
    "gia-starter-app-V1/internal/modules/auth/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Module struct {
    handler *http.AuthHandler
}

func NewModule(db *gorm.DB) *Module {
    handler := http.NewAuthHandler()

    return &Module{handler: handler}
}

func (m *Module) Register(r *gin.RouterGroup) {
    group := r.Group("/auth")
    group.GET("", m.handler.Index)
}

package cli

import "fmt"

func moduleTemplate(name string, pascalName string) string {
	return fmt.Sprintf(`package %s

import (
    "gia-starter-app-V1/internal/modules/%s/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Module struct {
    handler *http.%sHandler
}

func NewModule(db *gorm.DB) *Module {
    handler := http.New%sHandler()

    return &Module{handler: handler}
}

func (m *Module) Register(r *gin.RouterGroup) {
    group := r.Group("/%s")
    group.GET("", m.handler.Index)
}
`, name, name, pascalName, pascalName, name)
}

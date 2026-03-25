package users

import (
	"gia-starter-app-V1/internal/modules/users/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	handler *http.UsersHandler
}

func NewModule(db *gorm.DB) *Module {
	handler := http.NewUsersHandler()

	return &Module{handler: handler}
}

func (m *Module) Register(r *gin.RouterGroup) {
	group := r.Group("/users")
	group.GET("", m.handler.Index)
}

/*

	id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	username varchar(100) NOT NULL,
	email varchar(150) NOT NULL,
	"password" varchar(255) NOT NULL,
	id_role int4 NOT NULL,
	is_active bool DEFAULT true NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	deleted_at timestamp NULL,
*/

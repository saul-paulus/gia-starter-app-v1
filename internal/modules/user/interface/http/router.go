package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *UserHandler) {
	users := r.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("", h.GetAllUsers)
		users.GET("/:id", h.GetUserByID)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}

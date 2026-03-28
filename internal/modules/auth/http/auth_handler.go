package http

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Index godoc
// @Summary      Auth index
// @Description  Returns a welcome message from the auth module
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /auth [get]
func (h *AuthHandler) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello from auth module"})
}

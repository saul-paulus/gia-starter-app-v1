package http

import "github.com/gin-gonic/gin"

type AuthHandler struct{} // e.g. UserHandler

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{}
}

func (h *AuthHandler) Index(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Hello from auth module"})
}

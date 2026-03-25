package http

import "github.com/gin-gonic/gin"

type UsersHandler struct{} // e.g. UserHandler

func NewUsersHandler() *UsersHandler {
    return &UsersHandler{}
}

func (h *UsersHandler) Index(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Hello from users module"})
}

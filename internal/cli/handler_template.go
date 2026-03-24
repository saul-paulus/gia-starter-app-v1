package cli

import "fmt"

func handlerTemplate(name string, pascalName string) string {
	return fmt.Sprintf(`package http

import "github.com/gin-gonic/gin"

type %sHandler struct{} // e.g. UserHandler

func New%sHandler() *%sHandler {
    return &%sHandler{}
}

func (h *%sHandler) Index(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Hello from %s module"})
}
`, pascalName, pascalName, pascalName, pascalName, pascalName, name)
}
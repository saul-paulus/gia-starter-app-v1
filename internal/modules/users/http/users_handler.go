package http

import (
	"gia-starter-app-V1/internal/modules/users/dto"
	"gia-starter-app-V1/internal/modules/users/services"
	appErr "gia-starter-app-V1/internal/shared/errors"
	"gia-starter-app-V1/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	usersService services.UsersService
}

func NewUsersHandler(usersService services.UsersService) *UsersHandler {
	return &UsersHandler{usersService: usersService}
}

func (h *UsersHandler) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello from users module"})
}

func (h *UsersHandler) CreateUserHandler(c *gin.Context) {
	var reqUser dto.CreateUser

	// Binding + validasi request body
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		res := response.ApiErrorResponse(appErr.NewAppError(http.StatusBadRequest, "VALIDATION_ERROR", err.Error()))
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.usersService.CreateUser(reqUser); err != nil {
		res := response.ApiErrorResponse(err)
		if e, ok := err.(*appErr.AppError); ok {
			c.JSON(e.Status, res)
		} else {
			c.JSON(http.StatusInternalServerError, res)
		}
		return
	}

	res := response.ApiSuccessResponse(http.StatusCreated, "User created successfully", nil)
	c.JSON(http.StatusCreated, res)
}

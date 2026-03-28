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

// Index godoc
// @Summary      List users
// @Description  Returns a welcome message from the users module
// @Tags         Users
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /users [get]
func (h *UsersHandler) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello from users module"})
}

// CreateUserHandler godoc
// @Summary      Create a new user
// @Description  Creates a new user account. Email must be unique and the password will be hashed using bcrypt before storage.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUser  true  "Create user request payload"
// @Success      201      {object}  response.Response                            "User created successfully"
// @Failure      400      {object}  response.Response{error=map[string]string}  "Validation error or email already registered"
// @Failure      500      {object}  response.Response{error=map[string]string}  "Internal server error"
// @Router       /users [post]
func (h *UsersHandler) CreateUserHandler(c *gin.Context) {
	var reqUser dto.CreateUser

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

package handler

import (
	"gia-starter-app-V1/internal/domain/entity"
	"gia-starter-app-V1/internal/usecase"
	"gia-starter-app-V1/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUC: uc,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrorResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	if err := h.userUC.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, response.ApiSuccessResponse(http.StatusCreated, "User created successfully", user))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUC.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ApiSuccessResponse(http.StatusOK, "Users retrieved successfully", users))
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrorResponse(http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	user, err := h.userUC.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ApiErrorResponse(http.StatusNotFound, "User not found", nil))
		return
	}

	c.JSON(http.StatusOK, response.ApiSuccessResponse(http.StatusOK, "User retrieved successfully", user))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrorResponse(http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	var user entity.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrorResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}
	user.ID = id

	if err := h.userUC.UpdateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ApiSuccessResponse(http.StatusOK, "User updated successfully", user))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ApiErrorResponse(http.StatusBadRequest, "Invalid ID", nil))
		return
	}

	if err := h.userUC.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ApiSuccessResponse(http.StatusOK, "User deleted successfully", nil))
}

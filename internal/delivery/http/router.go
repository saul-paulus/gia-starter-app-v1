package http

import (
	_ "gia-starter-app-V1/docs"
	"gia-starter-app-V1/internal/delivery/http/middleware"
	"gia-starter-app-V1/internal/modules/user/delivery/http/handler"
	"gia-starter-app-V1/internal/shared/errors"
	"gia-starter-app-V1/pkg/response"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gia-starter-app-V1/internal/infrastructure/container"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all routes, middleware, and special handlers for the application.
func SetupRouter(r *gin.Engine, reg *container.Registry) {
	// Handle Method Not Allowed
	r.HandleMethodNotAllowed = true

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handle No Route (404)
	r.NoRoute(func(c *gin.Context) {
		res := response.ApiErrorResponse(http.StatusNotFound, "Resource not found", gin.H{
			"code": "NOT_FOUND",
		})
		c.JSON(http.StatusNotFound, res)
	})

	// Handle No Method (405)
	r.NoMethod(func(c *gin.Context) {
		res := response.ApiErrorResponse(http.StatusMethodNotAllowed, "Method not allowed", gin.H{
			"code": "METHOD_NOT_ALLOWED",
		})
		c.JSON(http.StatusMethodNotAllowed, res)
	})

	// Global Middleware
	r.Use(middleware.ErrorHandler())

	// API Routes
	v1 := r.Group("/api/v1")
	{
		// User Routes
		userHandler := handler.NewUserHandler(reg.UserUseCase)
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// @Summary      Health check
		// @Description  Check if the application is up and running
		// @Tags         system
		// @Produce      json
		// @Success      200  {object}  response.Response
		// @Router       /health [get]
		v1.GET("/health", func(c *gin.Context) {
			res := response.ApiSuccessResponse(http.StatusOK, "Health check OK", gin.H{
				"status": "UP OK",
			})
			c.JSON(http.StatusOK, res)
		})

		// @Summary      Error demo
		// @Description  Endpoint to simulate an error handled by middleware
		// @Tags         system
		// @Produce      json
		// @Success      404  {object}  response.Response
		// @Router       /error [get]
		v1.GET("/error", func(c *gin.Context) {
			// Simulating a "Not Found" error
			_ = c.Error(errors.ErrNotFound)
		})
	}
}

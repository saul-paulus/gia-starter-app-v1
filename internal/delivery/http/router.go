package http

import (
	_ "gia-starter-app-V1/docs"
	"gia-starter-app-V1/internal/modules/users"
	"gia-starter-app-V1/internal/shared/errors"
	"gia-starter-app-V1/internal/shared/middleware"
	"gia-starter-app-V1/internal/shared/response"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter configures all routes, middleware, and special handlers for the application.
func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// Handle Method Not Allowed
	r.HandleMethodNotAllowed = true

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handle No Route (404)
	r.NoRoute(func(c *gin.Context) {
		res := response.ApiErrorResponse(errors.ErrNotFound)
		c.JSON(http.StatusNotFound, res)
	})

	// Handle No Method (405)
	r.NoMethod(func(c *gin.Context) {
		res := response.ApiErrorResponse(errors.ErrBadRequest)
		c.JSON(http.StatusMethodNotAllowed, res)
	})

	// Global Middleware
	r.Use(middleware.ErrorHandler())

	// Initialize Modules
	usersModule := users.NewModule(db)

	// API Routes
	v1 := r.Group("/api/v1")
	{
		// @Summary      Health check
		// @Description  Memeriksa apakah aplikasi sedang berjalan dengan normal
		// @Tags         System
		// @Produce      json
		// @Success      200  {object}  response.Response
		// @Router       /health [get]
		v1.GET("/health", func(c *gin.Context) {
			res := response.ApiSuccessResponse(http.StatusOK, "Health check OK", gin.H{
				"status": "UP OK",
			})
			c.JSON(http.StatusOK, res)
		})

		// Register Users Module routes: GET /users, POST /users
		usersModule.Register(v1)
	}
}

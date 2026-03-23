package main

import (
	"gia-starter-app-V1/internal/delivery/http/middleware"
	"gia-starter-app-V1/internal/shared/errors"
	"gia-starter-app-V1/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Handle Method Not Allowed
	router.HandleMethodNotAllowed = true

	// Register Middleware
	router.Use(middleware.ErrorHandler())

	// Handle No Route (404)
	router.NoRoute(func(c *gin.Context) {
		res := response.ApiErrorResponse(http.StatusNotFound, "Resource not found", gin.H{
			"error": "NOT_FOUND",
		})
		c.JSON(http.StatusNotFound, res)
	})

	// Handle No Method (405)
	router.NoMethod(func(c *gin.Context) {
		res := response.ApiErrorResponse(http.StatusMethodNotAllowed, "Method not allowed", gin.H{
			"error": "METHOD_NOT_ALLOWED",
		})
		c.JSON(http.StatusMethodNotAllowed, res)
	})

	{
		v1 := router.Group("/api/v1")
		v1.GET("/health", func(c *gin.Context) {
			res := response.ApiSuccessResponse(http.StatusOK, "Health check OK", gin.H{
				"status": "UP",
			})
			c.JSON(http.StatusOK, res)
		})

		v1.GET("/error", func(c *gin.Context) {
			// Simulating a "Not Found" error using c.Error()
			_ = c.Error(errors.ErrNotFound)
		})
	}

	router.Run(":8081")
}

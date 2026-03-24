package middleware

import (
	stdErrors "errors"
	appErrors "gia-starter-app-V1/internal/shared/errors"
	"gia-starter-app-V1/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that catches errors set via c.Error().
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var appErr *appErrors.AppError

		if stdErrors.As(err, &appErr) {
			res := response.ApiErrorResponse(appErr.Status, appErr.Message, gin.H{
				"code": appErr.Code,
			})
			c.AbortWithStatusJSON(appErr.Status, res)
		} else {
			res := response.ApiErrorResponse(http.StatusInternalServerError, "An unexpected error occurred", gin.H{
				"code": "INTERNAL_SERVER_ERROR",
			})
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		}
	}
}

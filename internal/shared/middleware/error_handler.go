package middleware

import (
	stdErrors "errors"
	appErrors "gia-starter-app-V1/internal/shared/errors"
	"gia-starter-app-V1/internal/shared/response"
	"log"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that catches errors set via c.Error().

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		// ambil error terakhir (bisa di-upgrade nanti)
		err := c.Errors.Last().Err

		var appErr *appErrors.AppError

		if stdErrors.As(err, &appErr) {
			res := response.ApiErrorResponse(appErr)
			c.AbortWithStatusJSON(appErr.Status, res)
			return
		}

		// log error asli (penting untuk debugging)
		log.Println("Unhandled error:", err)

		// fallback ke internal error standar
		res := response.ApiErrorResponse(appErrors.ErrInternal)
		c.AbortWithStatusJSON(appErrors.ErrInternal.Status, res)
	}
}

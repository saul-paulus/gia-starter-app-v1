package main

// @title           GIA Starter App API
// @version         1.0
// @description     This is a starter kit for Go Gin Clean Architecture.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Saul Paulus
// @contact.url    https://github.com/saul-paulus

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8081
// @BasePath  /api/v1

import (
	"gia-starter-app-V1/internal/delivery/http"
	"gia-starter-app-V1/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Logger
	logger.InitLogger()
	logger.Info("Starting GIA Starter App...")

	router := gin.Default()

	// Setup Router
	http.SetupRouter(router)

	router.Run(":8081")
}

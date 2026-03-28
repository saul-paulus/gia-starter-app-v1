// Package main is the entry point for the GIA Starter App API.
//
// @title           GIA Starter App API
// @version         1.0
// @description     A professional-grade backend starter kit built with Golang and Gin, following Modular Clean Architecture.
//
// @contact.name    Saul Paulus
// @contact.url     https://github.com/saul-paulus/gia-starter-app-v1
//
// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT
//
// @host            localhost:8081
// @BasePath        /api/v1
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.
package main

import (
	"gia-starter-app-V1/internal/bootstrap"
)

func main() {
	bootstrap.InitApp()
}

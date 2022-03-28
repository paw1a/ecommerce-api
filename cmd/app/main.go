package main

import "github.com/paw1a/ecommerce-api/internal/app"

// @title        E-commerce API
// @version      1.0
// @description  This is simple api of e-commerce shop

// @contact.name   API Support
// @contact.url    https://t.me/paw1a
// @contact.email  paw1a@yandex.ru

// @host      52.29.184.51:8080
// @BasePath  /api/v1

// @schemes  http

// @securityDefinitions.apikey  AdminAuth
// @in                          header
// @name                        Authorization

// @securityDefinitions.apikey  UserAuth
// @in                          header
// @name                        Authorization
func main() {
	app.Run("config/config.yml")
}

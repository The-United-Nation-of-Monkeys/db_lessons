package main

import "github.com/The-United-Nation-of-Monkeys/db_lessons/internal/container"

// @title Online Classes API
// @version 1.0
// @description API для управления онлайн-курсами
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	container.NewApp()
}

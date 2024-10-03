package main

import (
	"User-Post-Backend/infra"
	"User-Post-Backend/infra/logger"
	"User-Post-Backend/internal/handlers"
	"User-Post-Backend/internal/middleware"
	"log"
	"os"

	_ "User-Post-Backend/docs"

	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

// @title UserPost API
// @version 1.0
// @description API documentation for UserPost backend.
// @host localhost:3000
// @BasePath /
func main() {
	logger.InitializeLogger("app.log")

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(middleware.CORS())
	app.Use(middleware.ErrorHandlerMiddleware)

	db := infra.InitDB()
	cache := infra.NewRedisClient()

	handlers.NewUserHandler(app, db, cache)
	handlers.NewPostHandler(app, db, cache)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Set default port ke 3000 jika PORT tidak ada
	}
	app.Listen(":" + port)
}

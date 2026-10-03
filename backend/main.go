package main

import (
	"holygrail-backend/config"
	"holygrail-backend/controllers"
	"holygrail-backend/middleware"
	"holygrail-backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// [STEP 1.1] Load the environment-backed settings needed by the HTTP server and account service.
	settings, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// [STEP 1.1] Initialize the authentication database and keep its handle alive for the server lifetime.
	database, err := config.Connect(settings)
	if err != nil {
		log.Fatal("Failed to connect to authentication database:", err)
	}
	defer database.Close()

	// [STEP 1.1] Boot the Gin application, middleware, handlers, and API routes.
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	handlers := controllers.New(database.Auth, settings)
	routes.SetupRoutes(r, handlers, middleware.Auth(database.Auth, settings.JWTSecret))

	log.Printf("Backend listening on :%s", settings.Port)
	if err := r.Run(":" + settings.Port); err != nil {
		log.Fatal("Server stopped:", err)
	}
}

package main

import (
	"log"
	"test/controllers/database"
	"test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env.
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Failed to load .env file: %v\n", err)
	}

	//calling database connection function
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")
	// Create a new Fiber instance something like server anouncemnet
	r := fiber.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	//make the route for the server
	routes.SetupRoutes(r, db)

	//this is the route for the server to listen on port 3000
	log.Fatal(r.Listen(":3030"))
}

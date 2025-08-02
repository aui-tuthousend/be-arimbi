package main

import (
	"be-arimbi/container"
	"be-arimbi/routes"
	"be-arimbi/internal/database"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()

	if os.Getenv("ENV") == "production" {
		log.Println("Running in production mode")
	} else {
		log.Println("Running in development mode")
	}
}

func main() {
	app := fiber.New()
	
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		// AllowOrigins: "http://localhost:3000",
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",         
	}))

	db := database.Connect()
	c := container.InitApp(db)
	routes.SetupRoutes(app, c)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("sup ni99a")
	})

	log.Fatal(app.Listen(":9000"))
}
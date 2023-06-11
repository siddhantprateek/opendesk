package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	connection "github.com/nikhitabarat/opendesk/connection"
	controllers "github.com/nikhitabarat/opendesk/handlers/controllers"
)

// @desp: Opendesk API Handlers
func OpenDeskApiHandler() {
	router := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	router.Use(cors.New())

	// Connect Database
	connection.MongoDBdatabase()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is Healthy")
	})

	router.Post("/task", controllers.AddTask)
	router.Get("/task/:taskId", controllers.GetTaskbyId)
	router.Get("/tasks", controllers.GetEmployeeTask)
	router.Delete("/task/:taskId", controllers.DeleteTaskbyId)

	router.Get("/quotes", controllers.GetAllQuotes)
	router.Post("/quote", controllers.CreateMotivationQuotes)
	router.Get("/quote/:quoteId", controllers.GetQuotesbyId)
	router.Delete("/quote/:quoteId", controllers.DeleteQuotes)

	router.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})
	router.Listen(":8000")

}

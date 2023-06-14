package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/siddhantprateek/opendesk/configs"
	handlers "github.com/siddhantprateek/opendesk/storage/handlers"
)

func storageAPIroutes(router *fiber.App) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Opendesk Storage API.",
		})
	})

	task := router.Group("/api/task")

	task.Post("/task", handlers.AddTask)
	task.Get("/task/:taskId", handlers.GetTaskbyId)
	task.Get("/tasks", handlers.GetEmployeeTask)
	task.Delete("/task/:taskId", handlers.DeleteTaskbyId)

	quote := router.Group("/qoute")

	quote.Get("/all", handlers.GetAllQuotes)
	quote.Post("/new", handlers.CreateMotivationQuotes)
	quote.Get("/:quoteId", handlers.GetQuotesbyId)
	quote.Delete("/:quoteId", handlers.DeleteQuotes)
	quote.Patch("/:quoteId", handlers.UpdateMotivationQuotes)
}

func Init() error {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Cross-Origin Resource Sharing
	app.Use(cors.New())

	// logger
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	storageAPIroutes(app)

	PORT := configs.GetEnv("STORAGE_PORT")
	err := app.Listen(":" + PORT)
	return err
}

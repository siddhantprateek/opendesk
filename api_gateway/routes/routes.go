package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handler "github.com/siddhantprateek/opendesk/api_gateway/handler"
	apiMiddleware "github.com/siddhantprateek/opendesk/api_gateway/middleware"
	"github.com/siddhantprateek/opendesk/configs"
)

func authRoutes(authr *echo.Group) {
	authr.GET("/", handler.AuthInit)
	authr.POST("/create", handler.CreateUser)
}

func gatewayRoutes(route *echo.Echo) {
	route.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Welcome to Opendesk API Gateway.",
		})
	})
}

func Init() error {

	app := echo.New()

	// Cross-Origin Resource Sharing.
	app.Use(middleware.CORS())
	// Logger
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// Key based API Authorization
	app.Use(middleware.KeyAuth(apiMiddleware.KeyAuthMiddleware))

	gatewayRoutes(app)

	authr := app.Group("/auth")
	authRoutes(authr)

	PORT := configs.GetEnv("API_GATEWAY_PORT")
	err := app.Start(":" + PORT)
	return err
}

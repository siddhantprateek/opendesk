package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/siddhantprateek/opendesk/configs"
)

// curl -H "Authorization: Bearer your-valid-key" http://localhost:8080/your-endpoint
func KeyAuthMiddleware(key string, c echo.Context) (bool, error) {
	return key == configs.GetEnv("API_AUTH_KEY"), nil
}

package routes

import (
	"testing"

	"holygrail-backend/controllers"

	"github.com/gin-gonic/gin"
)

func TestAuthEndpointsAreRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	SetupRoutes(router, &controllers.Handlers{}, func(c *gin.Context) { c.Next() })

	wanted := map[string]bool{
		"GET /":                   false,
		"POST /api/auth/register": false,
		"POST /api/auth/login":    false,
		"POST /api/auth/logout":   false,
	}
	for _, route := range router.Routes() {
		key := route.Method + " " + route.Path
		if _, exists := wanted[key]; exists {
			wanted[key] = true
		}
	}
	for route, found := range wanted {
		if !found {
			t.Errorf("%s is not registered", route)
		}
	}
}

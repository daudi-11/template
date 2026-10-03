package routes

import (
	"github.com/gin-gonic/gin"
	"holygrail-backend/controllers"
)

// [STEP 3.1] Register the public home and the account lifecycle endpoints in one place.
func SetupRoutes(router *gin.Engine, handlers *controllers.Handlers, auth gin.HandlerFunc) {
	router.GET("/", handlers.Home)

	api := router.Group("/api/auth")
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)
	api.POST("/logout", auth, handlers.Logout)
}

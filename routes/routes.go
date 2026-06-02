package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Events Endpoints
	events := server.Group("/events")

	events.Use(middlewares.Authorization)

	events.GET("/", GetAllEvents)

	events.GET("/:id", GetEventById)

	events.POST("/", CreateEvent)

	events.PUT("/:id", UpdateEvent)

	events.DELETE("/:id", DeleteEvent)

	// Registers Endpoints
	registerEvents := server.Group("/register-event/:eventId")

	registerEvents.Use(middlewares.Authorization)

	registerEvents.POST("/", registerEvent)

	registerEvents.DELETE("/", cancelRegisterEvent)

	// Auth Endpoints
	auth := server.Group("/auth")

	auth.POST("/register", register)

	auth.POST("/login", login)
}

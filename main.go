package main

import (
	"example.com/DB"
	"example.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	DB.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

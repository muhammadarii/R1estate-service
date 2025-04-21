package main

import (
	"r1estate-service/config"
	"r1estate-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
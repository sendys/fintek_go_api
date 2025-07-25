package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	//config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.Run(":8081")
}

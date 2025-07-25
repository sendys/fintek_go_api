package main

import (
	"backend/config"
	"backend/models"
	"backend/routes"
	"backend/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	
	// Auto-migrate database models
	if err := config.DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize upload directory
	if err := utils.InitUploadDir(); err != nil {
		log.Fatal("Failed to create upload directory: ", err)
	}

	r := gin.Default()

	// Set up routes
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.ProductRoutes(r)

	r.Run(":8081")
}
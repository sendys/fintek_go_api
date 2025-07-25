package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	// Public routes (no authentication required)
	public := r.Group("/products")
	{
		public.GET("/", controllers.GetAllProducts)           // Get all products with pagination and filtering
		public.GET("/:id", controllers.GetProductByID)       // Get single product by ID
		public.GET("/categories", controllers.GetProductCategories) // Get all categories
	}

	// Protected routes (authentication required)
	protected := r.Group("/products")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/", controllers.CreateProduct)              // Create new product
		protected.PUT("/:id", controllers.UpdateProduct)           // Update product
		protected.DELETE("/:id", controllers.DeleteProduct)        // Delete product
		protected.POST("/:id/image", controllers.UploadProductImage) // Upload product image
	}

	// Serve static files for uploaded images
	r.Static("/uploads", "./uploads")
}
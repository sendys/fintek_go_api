package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	protected := r.Group("/user")

	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/:user_id", controllers.GetProfile) // gunakan parameter user_id
		protected.GET("/all", controllers.GetAllUsers)
	}
}

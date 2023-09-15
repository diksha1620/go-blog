package routes

import (
	"github.com/dish1620/controller"
	"github.com/dish1620/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// authenticated := r.Group("/api")
	// authenticated.Use(middleware.AuthMiddleware())
	public := r.Group("/public")
	{
		public.POST("/login", controller.Login)
		public.POST("/register", controller.CreateUser)
	}
	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{

		protected.GET("/users", controller.GetAllUsers)
		protected.GET("/users/:id", controller.GetUser)
		protected.DELETE("/deleteusers/:id", controller.DeleteUser)
		protected.PATCH("/users/:id", controller.UpdateUser)
	}
	r.Run(":8080")
	return r
}

package routes

import (
	"r1estate-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r*gin.Engine) {
	api := r.Group("/api/v1")
	{
		user := api.Group("/users")
		{
			user.POST("/register", controllers.CreateUser)
			user.GET("/", controllers.GetAllUsers)
			user.GET("/:id", controllers.GetUserByID)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
			user.POST("/signin", controllers.SignIn)
		}
		
		role:= api.Group("/roles")
		{
			role.POST("/", controllers.CreateRole)
			role.GET("/", controllers.GetRoles)
			role.GET("/:id", controllers.GetRoleByID)
			role.PUT("/:id", controllers.UpdateRole)
			role.DELETE("/:id", controllers.DeleteRole)
		}
	}
}
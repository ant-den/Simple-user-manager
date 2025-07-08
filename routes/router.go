package routes

import (
	"GO_Project/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *controller.UserController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", c.Create)
			user.GET("/", c.GetAll)
			user.GET("/:id", c.GetByID)
			user.PUT("/:id", c.Update)
			user.DELETE("/:id", c.Delete)
		}
	}
	return r
}

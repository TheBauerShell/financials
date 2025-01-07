package routes

import (
	"github.com/gin-gonic/gin"
	"your_project/backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/projects", controllers.CreateProject)
		api.GET("/projects", controllers.GetProjects)
		api.PUT("/projects/:id", controllers.UpdateProject)
		api.DELETE("/projects/:id", controllers.DeleteProject)
	}
}
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"your_project_path/backend/models"
	"your_project_path/backend/services"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create project"})
		return
	}
	c.JSON(http.StatusCreated, project)
}

func GetProjects(c *gin.Context) {
	projects, err := services.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func UpdateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update project"})
		return
	}
	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteProject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete project"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
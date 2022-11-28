package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-otero390/race_war/database"
	"github.com/leonardo-otero390/race_war/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, nil)
}

func AllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

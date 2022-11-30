package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-otero390/race_war/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/health", controllers.HealthCheck)
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.Run()
}

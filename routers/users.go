package routers

import (
	"godc/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersUsers(router *gin.RouterGroup) {
	router.GET("/users", controllers.Users)
	router.POST("/users", controllers.UserAdd)
}

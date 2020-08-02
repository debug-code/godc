package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"godc/controllers"
)

func RoutersUsers(router *gin.RouterGroup) {
	fmt.Println("RoutersUsers")
	router.GET("/users", controllers.Users)
	router.POST("/users", controllers.UserAdd)
}

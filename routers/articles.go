package routers

import (
	"godc/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersArticles(router *gin.RouterGroup) {
	router.GET("/articles", controllers.Users)
	router.POST("/articles", controllers.UserAdd)
	router.DELETE("/articles", controllers.UserAdd)
	router.PUT("/articles", controllers.UserAdd)
}

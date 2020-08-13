package routers

import (
	"godc/controllers"

	"github.com/gin-gonic/gin"
)

// RoutersArticles router
func RoutersArticles(router *gin.RouterGroup) {
	router.GET("/articles/*id", controllers.Articles)
	router.POST("/articles", controllers.ArticleAdd)
	router.DELETE("/articles", controllers.UserAdd)
	router.PUT("/articles", controllers.UserAdd)
}

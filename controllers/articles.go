package controllers

import (
	"fmt"
	"godc/common/vo"
	"godc/models"
	"godc/service"

	"github.com/gin-gonic/gin"
)

// Articles list
func Articles(this *gin.Context) {
	uid := this.Param("id")
	if uid != "" {

		article, err := service.ArticlesOne(uid)
		if err != nil {
			this.JSON(200, map[string]string{"error": err.Error()})
			return
		}
		vo.Success(this, article)
		return
	}
	// page := tool.PageSize(
	// 	this.Query("pageNum"),
	// 	this.Query("pageSize"))

	vo.Success(this, "OK")
	return
}

// ArticleAdd add a article
func ArticleAdd(this *gin.Context) {

	articleDto := models.Article{}
	err := this.ShouldBindJSON(&articleDto)
	if err != nil {
		this.JSON(200, map[string]string{"error": err.Error()})
		return
	}
	fmt.Println(articleDto)
	err = service.ArticlesAdd(articleDto)
	if err != nil {
		this.JSON(200, map[string]string{"error": err.Error()})
		return
	}

	vo.Success(this, articleDto)
	return
}

// ArticleModify  modify article
func ArticleModify(this *gin.Context) {
	vo.Success(this, "OK")
	return
}

// ArticleDelete  modify article
func ArticleDelete(this *gin.Context) {
	vo.Success(this, "OK")
	return
}

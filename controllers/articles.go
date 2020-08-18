package controllers

import (
	"fmt"
	"godc/common/tool"
	"godc/common/vo"
	"godc/models"
	"godc/service"
	"strings"

	"debug-code.com/dclog"

	"github.com/gin-gonic/gin"
)

// Articles list
func Articles(this *gin.Context) {
	uid := this.Param("id")
	dclog.Info(uid)
	if uid != "" {

		article, err := service.ArticleOne(uid)
		if err != nil {
			dclog.Error(err)
			this.JSON(200, map[string]string{"error": err.Error()})
			return
		}
		vo.Success(this, article)
		return
	}
	page := tool.PageSize(
		this.Query("pageNum"),
		this.Query("pageSize"))

	article := models.Article{}

	res, err := service.Articles(article, page)
	if err != nil {
		dclog.Error(err)
		this.JSON(200, map[string]string{"error": err.Error()})
		return
	}

	vo.Success(this, res)
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

// ArticleDelete  delete  article
func ArticleDelete(this *gin.Context) {
	flag := this.Query("flag")
	ids := this.Query("ids")
	if ids == "" {
		vo.Success(this, "Fail")
		return
	}
	flagBool := false
	if flag != "" {
		flagBool = true
	}
	arr := strings.Split(ids, ",")

	err := service.ArticlesDel(arr, flagBool)
	if err != nil {
		dclog.Error(err)
		vo.Success(this, "Fail")
		return
	}
	vo.Success(this, "OK")
	return

}

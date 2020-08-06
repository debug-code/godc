package controllers

import (
	"fmt"
	"godc/common/vo"
	"godc/service"

	"github.com/gin-gonic/gin"
)

// Articles list
func Articles(ctx *gin.Context) {
	articles, err := service.Articles()
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
	}

	// fmt.Println(users)
	vo.Success(ctx, articles)
	return
}

// ArticleAdd add a article
func ArticleAdd(ctx *gin.Context) {
	users, err := service.Users()
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
	}

	fmt.Println(users)
	// vo.Success(ctx, users)
	return
}

// ArticleModify  modify article
func ArticleModify(ctx *gin.Context) {
	users, err := service.Users()
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
	}

	fmt.Println(users)
	// vo.Success(ctx, users)
	return
}

// ArticleDelete  modify article
func ArticleDelete(ctx *gin.Context) {
	users, err := service.Users()
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
	}

	fmt.Println(users)
	// vo.Success(ctx, users)
	return
}

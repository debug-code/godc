package controllers

import (
	"fmt"
	"godc/models"
	"godc/service"

	"godc/common/vo"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Users(ctx *gin.Context) {
	users, err := service.Users()
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
	}

	fmt.Println(users)
	vo.Success(ctx, users)

	return
}

func UserAdd(ctx *gin.Context) {
	fmt.Print("UserAdd")
	db, err := gorm.Open("dbMysql",
		"alex:123qwe=-0@(localhost:33306)/godc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()

	user := models.User{Uid: "sdsf", UName: "sdf"}
	sd := db.Create(&user)
	if sd.Error != nil {
		fmt.Println("POST", sd.Error.Error())
	}

	ctx.JSON(200, user)
	return
}

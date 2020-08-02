package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"godc/models"
)

func Users(ctx *gin.Context) {
	fmt.Println("Users")
	db, err := gorm.Open("mysql",
		"alex:123qwe=-0@(localhost:33306)/godc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()

	users := []models.User{}
	db.Find(&users)
	//sd :=
	//fmt.Println("GET", sd.Error.Error())
	fmt.Println(users)
	ctx.JSON(200, users)
	return
}

func UserAdd(ctx *gin.Context) {
	fmt.Print("UserAdd")
	db, err := gorm.Open("mysql",
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

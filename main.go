package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"godc/routers"
)

func init() {
	db, err := gorm.Open("mariadb",
		"alex:123qwe=-0@/controllers?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()
}
func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	routers.Router(engine)
	// 绑定端口，然后启动应用
	engine.Run(":9527")
}

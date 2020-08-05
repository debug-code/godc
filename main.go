package main

import (
	"github.com/gin-gonic/gin"
	"godc/db"
	"godc/routers"
)

func init() {
	//初始化数据库
	db.DbInit()
}
func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	routers.Router(engine)
	// 绑定端口，然后启动应用
	engine.Run(":9527")
}

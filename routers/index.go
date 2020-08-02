package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	fmt.Println("Router")
	v1 := router.Group("/v1")
	RoutersUsers(v1)
}

package vo

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, res interface{}) {
	ctx.JSON(200, res)
}

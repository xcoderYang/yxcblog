package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ParamsError(ctx *gin.Context, err error){
	log.Println(err)
	ctx.JSON(200, gin.H{
		"msg": "查询参数有误",
		"success": false,
	})
}

func ServerError(ctx *gin.Context, err error){
	log.Println(err)
	ctx.JSON(200, gin.H{
		"msg": "服务器处理错误",
		"success": false,
	})
}

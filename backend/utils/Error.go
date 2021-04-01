package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ParamsError(ctx *gin.Context, err error, str... string){
	log.Println(err)
	var res string
	if len(str) == 0{
		res = "查询参数错误"
	}else{
		res = str[0]
	}
	ctx.JSON(200, gin.H{
		"msg": res,
		"success": false,
	})
}

func ServerError(ctx *gin.Context, err error, str... string){
	log.Println(err)
	var res string
	if len(str) == 0{
		res = "服务器处理错误"
	}else{
		res = str[0]
	}
	ctx.JSON(500, gin.H{
		"msg": res,
		"success": false,
	})
}
func SessionError(ctx *gin.Context, err error, str... string){
	log.Println(err)
	var res string
	if len(str) == 0{
		res = "服务器处理错误"
	}else{
		res = str[0]
	}
	ctx.JSON(500, gin.H{
		"msg": res,
		"success": false,
	})
}

package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"yxcblog/dbInit"
)


var (
	DB  = dbInit.DB
	REDIS = dbInit.REDIS
)

func CheckLogin() gin.HandlerFunc{
	return func(ctx *gin.Context){
		sessionId, err := ctx.Cookie("SessionId")
		contx := context.Background()
		tag := ""
		if err!=nil{
			fmt.Println(err)
			tag = "SessionID 查询出错"
		}else{
			if sessionId == ""{
				tag = "SessionID已过期"
			}else{
				test, _ := REDIS.Get(contx, sessionId).Result()
				if test == ""{
					tag = "SessionID已过期"
				}
			}
		}
		if tag != ""{
			ctx.JSON(401, gin.H{
				"msg": tag,
				"success": false,
			})
			//utils.SessionError(ctx, err, tag)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

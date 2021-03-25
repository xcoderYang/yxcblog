package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"yxcblog/dbInit"
	"yxcblog/utils"
)


var (
	DB  = dbInit.DB
	REDIS = dbInit.REDIS
)

func CheckLogin() gin.HandlerFunc{
	return func(ctx *gin.Context){
		sessionId, err := ctx.Cookie("SessionID")
		contx := context.Background()
		if err!=nil{
			utils.SessionError(ctx, "SessionID 查询出错", err)
		}else if sessionId == ""{
			utils.SessionError(ctx, "SessionID已过期", err)
		}else{
			//test, _ := REDIS.Get(contx, "test").Result()
			//fmt.Print(test)
		}
		fmt.Println("start check")
		test, err := REDIS.Get(contx, "test").Result()
		if err != nil{
			fmt.Println("redis check error")
			fmt.Println(err)
		}
		fmt.Println(test)
		ctx.Next()
	}
}

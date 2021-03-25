package middleware

import (
	"github.com/gin-gonic/gin"
	"yxcblog/utils"
)


//var DB  = database.DB

func CheckLogin() gin.HandlerFunc{
	return func(ctx *gin.Context){
		sessionId, err := ctx.Cookie("SessionID")
		if err!=nil{
			utils.SessionError(ctx, err)
		}else if sessionId == ""{

		}
		ctx.Next()
	}
}

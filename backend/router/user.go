package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"yxcblog/global"
	"log"
	"strconv"
	"time"
	"yxcblog/dao"
	"yxcblog/dbInit"
	_ "yxcblog/middleware"
	"yxcblog/utils"
)

func initUser(r *gin.RouterGroup){
	v1:=r.Group("/user")
	{
		v1.GET("/registry", registry)
		v1.GET("/login", login)
	}
}

func registry(ctx *gin.Context){
	log.Println("addUser")
	username := ctx.Query("username")
	password := ctx.Query("password")
	email := ctx.Query("email")

	var user dao.User
	user.Email = email
	user.Username = username
	user.Password = fmt.Sprintf("%x", utils.Sha256(password))
	user.AddTime = strconv.FormatInt(time.Now().UnixNano(), 10)
	user.Introducer = 0
	user.LastLogin = user.AddTime
	user.LabelId = "default"
	user.Mobile = ""
	err := dao.AddUser(user)
	if err!=nil{
		utils.ParamsError(ctx, err, "参数设置错误")
		return
	}
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
		"success": true,
	})
}
func login(ctx *gin.Context){
	username := ctx.Query("username")
	password := ctx.Query("password")
	userInfo, err := dao.QueryByUserUserName(username, password)
	log.Println("userInfo is ", userInfo)
	if err!=nil{
		utils.ParamsError(ctx, err, "登录信息有误")
		return
	}
	REDIS := dbInit.REDIS
	c := context.Background()
	ans, err := REDIS.Set(c, username, userInfo, global.REDIS_EXPIRE_TIME).Result()
	if err != nil{
		utils.ServerError(ctx, err, "服务器错误")
		return
	}
	log.Println(ans)
	ctx.JSON(200, gin.H{
		"msg": "登录成功",
		"success": true,
	})
}
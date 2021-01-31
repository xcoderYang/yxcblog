package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

type User struct{
	ID int
	AddTime string `gorm:"add_time"`
	Power int `gorm:"power"`
	Username string
	Nickname string
	Password string
	LabelId string `gorm:"label_id"`
	Mobile string
	Father int
	Email string
	LastLogin time.Time
}

//var DB  = database.DB

func CheckLogin() gin.HandlerFunc{
	return func(ctx *gin.Context){
		ctx.Next()
		//ans, err := ctx.Cookie("loginInfo")
		//if err!=nil{
		//	fmt.Println("error1")
		//}
		//if ans == ""{
		//	ctx.JSON(200, gin.H{
		//		"code": "F001",
		//		"message": "登录状态失效，请重新登录",
		//	})
		//	return
		//}
		//user := &User{
		//
		//}
		//DB.Where("password = ?", ans).First(user)
		//
		//fmt.Println(user)
		//
		//if user.Nickname==""{
		//	ctx.JSON(200, gin.H{
		//		"code": "F002",
		//		"message": "登录信息有误，请重新登录",
		//	})
		//	return
		//}
		//fmt.Println("login success")
		//ctx.JSON(200, gin.H{
		//	"code": "S001",
		//	"message": "登录成功",
		//})
		//return
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"time"
	_ "yxcblog/middleware"
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

func initUser(r *gin.RouterGroup){
	//dbPool := db.DB
	//var usr User
	//r.GET("/setcookie", func(ctx *gin.Context){
	//	ctx.SetCookie("loginInfo", "d49f513c1144fcfe319ada237299770743e018f043f69eee4556028ee0a619a8", 3600, "/", "localhost", false, true)
	//})
	//r.GET("/getcookie", func(ctx *gin.Context){
	//	cookie, err := ctx.Cookie("cookieTest")
	//	if err != nil{
	//		fmt.Println("no cookie")
	//	}else{
	//		fmt.Println("cookie is ", cookie)
	//	}
	//})
	//v1 := r.Group("/user")
	//{
	//	v1.GET("/login", func(ctx *gin.Context){
	//		account := ctx.Params.ByName("account")
	//		if account == "" {
	//			ctx.JSON(200, gin.H{
	//				"success": true,
	//				"msg":     "account is invalid",
	//			})
	//		}
	//		dbPool.First(&usr)
	//		fmt.Printf("%#v", usr)
	//	})
	//	v1.GET("logout", func(ctx *gin.Context){
	//
	//	})
	//}
}
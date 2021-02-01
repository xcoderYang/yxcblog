package router

import (
	"github.com/gin-gonic/gin"
	_ "yxcblog/database"
)

func initComment(r *gin.RouterGroup){
	v1 := r.Group("/comment")
	{
		v1.GET("/list", getComment)
	}
}

func getComment(ctx *gin.Context){

}
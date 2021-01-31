package router

import (
	"github.com/gin-gonic/gin"
	_ "yxcblog/database"
)

func initComment(r *gin.Engine){
	v1 := r.Group("/comment")
	{
		v1.GET("/list", func(ctx *gin.Context) {

		})
	}
}
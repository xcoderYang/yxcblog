package router

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func LoadRouter(r *gin.Engine){
	initUser(r)
	initBlog(r)
	//r.GET("/ping", func(c *gin.Context){
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
}
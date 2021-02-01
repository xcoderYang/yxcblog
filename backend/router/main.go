package router

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func LoadRouter(r *gin.Engine){
	group := r.Group("/v1")
	initUser(group)
	initBlog(group)
	//r.GET("/ping", func(c *gin.Context){
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
}
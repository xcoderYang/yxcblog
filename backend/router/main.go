package router

import (
	"github.com/gin-gonic/gin"
	"yxcblog/middleware"
)

var router *gin.Engine

func LoadRouter(r *gin.Engine){
	group := r.Group("/v1")
	group.Use(middleware.CheckLogin())
	initUser(group)
	initBlog(group)
	//r.GET("/ping", func(c *gin.Context){
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
}
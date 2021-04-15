package router

import (
	"github.com/gin-gonic/gin"
	"yxcblog/middleware"
)

var router *gin.Engine

func LoadRouter(r *gin.Engine){
	backApi := r.Group("/api")
	user := backApi.Group("/user")
	blog := backApi.Group("/blog")
	blog.Use(middleware.CheckLogin())
	initUser(user)
	initBlog(blog)

	frontApi := r.Group("/server")
	fblog := frontApi.Group("/blog")
	initBlog(fblog)
}
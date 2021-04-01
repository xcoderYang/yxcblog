package router

import (
	"github.com/gin-gonic/gin"
	"yxcblog/middleware"
)

var router *gin.Engine

func LoadRouter(r *gin.Engine){
	group := r.Group("/api")
	user := group.Group("/user")
	blog := group.Group("/blog")
	blog.Use(middleware.CheckLogin())
	initUser(user)
	initBlog(blog)
}
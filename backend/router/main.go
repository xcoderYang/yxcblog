package router

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func LoadRouter(r *gin.Engine){
	group := r.Group("/api")
	//group.Use(middleware.CheckLogin())
	initUser(group)
	initBlog(group)
}
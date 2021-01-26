package main

import "github.com/gin-gonic/gin"

import _ "yxcblog/router"

func main(){
	r := gin.Default()
	//router.LoadRouter(r)
	r.Run(":8080")
}
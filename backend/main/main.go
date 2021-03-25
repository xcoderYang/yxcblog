package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "yxcblog/config"
	"yxcblog/router"
)

func test() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("start")
		fmt.Println(c.Params)
		c.Next()
		fmt.Println("over")
	}
}

func main(){
	r := gin.Default()
	router.LoadRouter(r)
	r.Run(":8080")
}
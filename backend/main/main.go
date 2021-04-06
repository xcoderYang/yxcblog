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

type JSONForm struct{
	Id float64	`json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	VisitedN float64 `json:"visitedN"`
	CommentN float64 `json:"commentN"`
	Type []string `json:"type"`
	Label []string `json:"label"`
}


type BlogForm struct{
	blogForm JSONForm
}

func main(){
	r := gin.Default()
	router.LoadRouter(r)
	r.Run(":8081")
}
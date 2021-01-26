package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init(){
	fmt.Println(router)
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

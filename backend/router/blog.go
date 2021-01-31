package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
	database "yxcblog/database"
)

type blog struct{
	ID int
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Article
}

type Article struct{
	Title string
	Content string
	WorkNum int
	Type string
	Order int
	Tag string
	Visited int
}

type Comment struct{
	ID int
	Author int
	Time time.Time
	ReplyTo int
	Blog int
	BlackList bool
}

func initBlog(r *gin.Engine){
	DB := database.DB
	v1 := r.Group("/blog")
	{
		v1.GET("/list", func(ctx *gin.Context) {
			pageNum, err := strconv.Atoi(ctx.Query("pageNum"))
			if err!=nil{
				log.Println(err)
				return
			}
			rows, err := DB.Query("SELECT createAt, updateAt, title, content, `type`, `order`, tag, visited FROM blog ORDER BY createAt LIMIT ?,?;", (pageNum-1)*10, pageNum*10)
			if err != nil{
				log.Println(err)
				return
			}
			ans := make([]gin.H, 0)
			for i:=0;rows.Next();i++{
				var row blog
				err := rows.Scan(&row.CreateAt, &row.UpdateAt, &row.Title, &row.Content, &row.Type, &row.Order, &row.Tag, &row.Visited)
				if err!=nil{
					fmt.Println(err)
				}
				ans = append(ans, gin.H{
					"createAt":row.CreateAt,
					"updateAt":row.UpdateAt,
					"title": row.Title,
					"content": row.Content,
					"type": row.Type,
					"order": row.Order,
					"tag": row.Tag,
					"visited": row.Visited,
				})
			}
			ctx.JSON(200, gin.H{
				"success":true,
				"data": ans,
			})
		})
	}
}
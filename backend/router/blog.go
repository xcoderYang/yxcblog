package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
	"yxcblog/init"
	"yxcblog/utils"
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

var DB = init.DB

func initBlog(r *gin.RouterGroup){
	v1 := r.Group("/blog")
	{
		v1.GET("/list", getBlog)
		v1.GET("/blogItem", getBlogByItem)
	}
}

func getBlog(ctx *gin.Context) {
	pageNumStr := ctx.Query("pageNum")
	if pageNumStr == ""{
		pageNumStr = "1"
	}
	pageNum, err := strconv.Atoi(pageNumStr)
	if err!=nil{
		log.Println(err)
		return
	}
	rows, err := DB.Query("SELECT createAt, updateAt, title, `type`, `order`, tag, visited FROM blog ORDER BY createAt LIMIT ?,?;", (pageNum-1)*10, pageNum*10)
	if err != nil{
		utils.ParamsError(ctx, err)
		return
	}
	ans := make([]gin.H, 0)
	for rows.Next(){
		var row blog
		err := rows.Scan(&row.CreateAt, &row.UpdateAt, &row.Title, &row.Content, &row.Type, &row.Order, &row.Tag, &row.Visited)
		if err!=nil{
			utils.ParamsError(ctx, err)
			return
		}
		ans = append(ans, gin.H{
			"createAt":row.CreateAt,
			"updateAt":row.UpdateAt,
			"title": row.Title,
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
}

func getBlogByItem(ctx *gin.Context){
	blogId := ctx.Query("blogId")
	if blogId == ""{
		ctx.JSON(200, gin.H{
			"success": false,
			"data": []int{},
			"msg": "查询参数错误",
		})
	}
	rows, err := DB.Query("SELECT createAt, updateAt, title, content, `type`, `order`, tag, visited FROM blog WHERE blogId = ?;", blogId)
	if err != nil{
		utils.ParamsError(ctx, err)
		return
	}
	var ans gin.H
	for rows.Next(){
		var row blog
		err := rows.Scan(&row.CreateAt, &row.UpdateAt, &row.Title,&row.Content, &row.Type, &row.Order, &row.Tag, &row.Visited)
		if err!=nil{
			utils.ParamsError(ctx, err)
			return
		}
		ans=gin.H{
			"createAt":row.CreateAt,
			"updateAt":row.UpdateAt,
			"title": row.Title,
			"Content": row.Content,
			"type": row.Type,
			"order": row.Order,
			"tag": row.Tag,
			"visited": row.Visited,
		}
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"data": ans,
	})

}

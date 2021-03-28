package router

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yxcblog/dao"
	"yxcblog/dbInit"
	"yxcblog/utils"
)

var DB = dbInit.DB

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
		utils.ParamsError(ctx, nil)
		return
	}
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil{
		utils.ParamsError(ctx, nil)
		return
	}
	ans, err := dao.ReadBlogByPage(pageNum, 10)
	if err != nil{
		utils.ServerError(ctx, nil)
		return
	}
	ctx.JSON(200, gin.H{
		"success":true,
		"data": ans,
	})
}

func getBlogByItem(ctx *gin.Context){
	blogId := ctx.Query("blogId")
	if blogId == ""{
		utils.ParamsError(ctx, nil)
		return
	}
	ans, err := dao.ReadBlogById(blogId)
	if err != nil{
		utils.ServerError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"data": ans,
	})

}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"strconv"
	"yxcblog/dao"
	"yxcblog/dbInit"
	"yxcblog/utils"
)

var DB = dbInit.DB

func initBlog(r *gin.RouterGroup){
	r.GET("/list", getBlog)
	r.GET("/blogItem", getBlogByItem)
	r.POST("/updateBlogById", UpdateBlogById)
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


type BlogForm struct{
	Id int64	`json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	VisitedN int64 `json:"visitedN"`
	CommentN int64 `json:"commentN"`
	Type []string `json:"type"`
	Label []string `json:"label"`
}


func UpdateBlogById(ctx *gin.Context){
	blogForm,_ := ioutil.ReadAll(ctx.Request.Body)
	form := BlogForm{}
	jsonParse := gjson.ParseBytes(blogForm)
	form.Id = jsonParse.Get("blogForm.id").Int()
	form.Title = jsonParse.Get("blogForm.title").Value().(string)
	form.Content = jsonParse.Get("blogForm.content").String()
	form.VisitedN = jsonParse.Get("blogForm.visitedN").Int()
	form.CommentN = jsonParse.Get("blogForm.commentN").Int()
	form.Type = make([]string, 0)
	for _,v:=range jsonParse.Get("blogForm.type").Array(){
		form.Type = append(form.Type, v.String())
	}
	form.Label = make([]string, 0)
	for _,v:=range jsonParse.Get("blogForm.label").Array(){
		form.Label = append(form.Label, v.String())
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"data": "OK",
	})
}


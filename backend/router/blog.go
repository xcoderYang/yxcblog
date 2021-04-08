package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"strconv"
	"yxcblog/dao"
	"yxcblog/dbInit"
	"yxcblog/model/request"
	"yxcblog/utils"
)

var DB = dbInit.DB

func initBlog(r *gin.RouterGroup){
	r.GET("/list", getBlog)
	r.GET("/blogItem", getBlogByItem)
	r.POST("/updateBlogById", UpdateBlogById)
	r.POST("/createBlog", CreateBlog)
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
	log.Printf("%#v", ans)
	if err != nil{
		utils.ServerError(ctx, err, "服务器错误")
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

func UpdateBlogById(ctx *gin.Context){
	blogForm,_ := ioutil.ReadAll(ctx.Request.Body)
	form := request.BlogForm{}
	jsonParse := gjson.ParseBytes(blogForm)
	form.BlogId = jsonParse.Get("blogForm.blogId").String()
	form.Id = jsonParse.Get("blogForm.id").String()
	form.Title = jsonParse.Get("blogForm.title").String()
	form.Content = jsonParse.Get("blogForm.content").String()
	form.VisitedN = jsonParse.Get("blogForm.visitedN").String()
	form.CommentN = jsonParse.Get("blogForm.commentN").String()
	form.Type = ""
	for i,v:=range jsonParse.Get("blogForm.type").Array(){
		if i == 0{
			form.Type+=v.String()
		}else{
			form.Type+="|"+v.String()
		}
	}
	form.Label = ""
	for i,v:=range jsonParse.Get("blogForm.label").Array(){
		if i == 0{
			form.Label+=v.String()
		}else{
			form.Label+="|"+v.String()
		}
	}
	err := dao.UpdateBlogById(form.BlogId, form)
	if err !=  nil{
		utils.ParamsError(ctx, err, "参数设置错误")
		return
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"data": "OK",
	})
}

func CreateBlog(ctx *gin.Context){
	blogAddForm,_ := ioutil.ReadAll(ctx.Request.Body)
	form := request.BlogForm{}
	jsonParse := gjson.ParseBytes(blogAddForm)
	form.BlogId = jsonParse.Get("blogAddForm.blogId").String()
	form.Id = jsonParse.Get("blogAddForm.id").String()
	form.Title = jsonParse.Get("blogAddForm.title").String()
	form.Content = jsonParse.Get("blogAddForm.content").String()
	form.VisitedN = jsonParse.Get("blogAddForm.visitedN").String()
	form.CommentN = jsonParse.Get("blogAddForm.commentN").String()
	form.Type = jsonParse.Get("blogAddForm.type").String()
	form.Label = jsonParse.Get("blogAddForm.label").String()

	err := dao.CreateBlog(form)
	if err !=  nil{
		utils.ParamsError(ctx, err, "参数设置错误")
		return
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"data": "OK",
	})
}


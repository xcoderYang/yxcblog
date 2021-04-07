package dao

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"yxcblog/dbInit"
	"yxcblog/model"
	"yxcblog/model/request"
)

func CreateBlog(){}
func CreateComment(blogId string, fatherCommentId string){
	log.Println(blogId, fatherCommentId)
}

func ReadBlogById(blogId string)(gin.H, error){
	DB := dbInit.DB
	var blog model.Blog
	row := DB.QueryRow("SELECT createAt,updateAt,deleteAt, title,`content`, `type`, label, visitedN,blogId, commentN FROM blog WHERE blogId=?;", blogId)
	err := row.Scan(&blog.CreateAt, &blog.UpdateAt, &blog.DeleteAt, &blog.Title, &blog.Content, &blog.Type, &blog.Label, &blog.VisitedN, &blog.BlogId, &blog.CommentN)
	if err != nil{
		return nil, err
	}
	ans := gin.H{
		"createAt":blog.CreateAt,
		"updateAt":blog.UpdateAt,
		"deleteAt":blog.DeleteAt,
		"title": blog.Title,
		"content":blog.Content,
		"type": blog.Type,
		"label": blog.Label,
		"visitedN": blog.VisitedN,
		"blogId":blog.BlogId,
		"commentN": blog.CommentN,
	}
	return ans, nil
}
func ReadBlogByPage(pageNum int, pageSize int)([]gin.H, error){
	DB := dbInit.DB
	ans := make([]gin.H, 0)
	rows, err := DB.Query("SELECT `id`, createAt, updateAt,deleteAt, `title`,`content`,`type`, `label`, visitedN,blogId,commentN FROM blog ORDER BY createAt LIMIT ?,?;", (pageNum-1)*pageSize, pageNum*pageSize)
	if err != nil{
		return ans, err
	}
	for rows.Next(){
		var row model.Blog
		err := rows.Scan(&row.ID, &row.CreateAt, &row.UpdateAt, &row.DeleteAt, &row.Title, &row.Content, &row.Type, &row.Label, &row.VisitedN, &row.BlogId, &row.CommentN)
		if err!=nil{
			return ans, err
		}
		ans = append(ans, gin.H{
			"id": row.ID,
			"createAt":row.CreateAt,
			"updateAt":row.UpdateAt,
			"deleteAt":row.DeleteAt,
			"title": row.Title,
			"content":row.Content,
			"type": row.Type,
			"label": row.Label,
			"visitedN": row.VisitedN,
			"blogId":row.BlogId,
			"commentN": row.CommentN,
		})
	}
	return ans, nil
}
func ReadBlogCount()(int,error){
	DB := dbInit.DB
	count := 0
	row := DB.QueryRow("SELECT COUNT(*) FROM blog")
	err := row.Scan(&count)
	if err != nil{
		return -1, err
	}else{
		return count, err
	}
}
func ReadBlogByKeyValue(key string, value string)(int,error){
	DB := dbInit.DB
	count := 0
	row := DB.QueryRow("SELECT COUNT(*) FROM blog WHERE ?=?;", key, value)
	err := row.Scan(&count)
	if err != nil{
		return -1, err
	}else{
		return count, err
	}
}
func ReadCommentByBlogId(){}

func UpdateBlogById(blogId string, form request.BlogForm)error{
	DB := dbInit.DB
	log.Println(blogId, form)
	result, err := DB.Exec("UPDATE blog SET `title`=?,`content`=?,`visitedN`=?,`commentN`=?,`type`=?,`label`=? WHERE `blogId`=?", form.Title, form.Content, form.VisitedN, form.CommentN, form.Type, form.Label, blogId)
	fmt.Println("error1")
	if err != nil{
		return err
	}
	rows, err := result.RowsAffected()
	fmt.Println("error2")
	if err != nil{
		return err
	}
	fmt.Println("error3")
	if rows == 1{
		return nil
	}else{
		return errors.New("")
	}
}

func DeleteBlogById(blogId string){
	log.Println(blogId)
}
func DeleteComment(blogId string, commentId string){
	log.Println(blogId, commentId)
}

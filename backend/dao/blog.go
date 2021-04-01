package dao

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"yxcblog/dbInit"
)

type Blog struct{
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
	BlogId string
}

type Comment struct{
	ID int
	Author int
	Time time.Time
	ReplyTo int
	Blog int
	BlackList bool
}

func (b *Blog)Create(){

}
func (b *Blog)Delete(){

}
func (b *Blog)Update(){

}

func CreateBlog(){}
func CreateComment(blogId string, fatherCommentId string){
	log.Println(blogId, fatherCommentId)
}

func ReadBlogById(blogId string)(gin.H, error){
	DB := dbInit.DB
	var blog Blog
	row := DB.QueryRow("SELECT createAt, updateAt,deleteAt, title,`content`,workNum, `type`, `order`, tag, visited,blogId FROM blog WHERE blogId=?;", blogId)
	err := row.Scan(&blog.CreateAt, &blog.UpdateAt, &blog.DeleteAt, &blog.Title, &blog.Content, &blog.Type, &blog.Order, &blog.Tag, &blog.Visited, &blog.BlogId)
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
		"order": blog.Order,
		"tag": blog.Tag,
		"visited": blog.Visited,
		"blogId":blog.BlogId,
	}
	return ans, nil
}
func ReadBlogByPage(pageNum int, pageSize int)([]gin.H, error){
	DB := dbInit.DB
	ans := make([]gin.H, 0)
	rows, err := DB.Query("SELECT createAt, updateAt,deleteAt, title,`content`,workNum, `type`, `order`, tag, visited,blogId FROM blog ORDER BY createAt LIMIT ?,?;", (pageNum-1)*pageSize, pageNum*pageSize)
	if err != nil{
		return ans, err
	}
	for rows.Next(){
		var row Blog
		err := rows.Scan(&row.CreateAt, &row.UpdateAt, &row.DeleteAt, &row.Title, &row.Content, &row.Type, &row.Order, &row.Tag, &row.Visited, &row.BlogId)
		if err!=nil{
			return ans, err
		}
		ans = append(ans, gin.H{
			"createAt":row.CreateAt,
			"updateAt":row.UpdateAt,
			"deleteAt":row.DeleteAt,
			"title": row.Title,
			"content":row.Content,
			"type": row.Type,
			"order": row.Order,
			"tag": row.Tag,
			"visited": row.Visited,
			"blogId":row.BlogId,
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

func UpdateBlogById(blogId string) {
	log.Println(blogId)
}

func DeleteBlogById(blogId string){
	log.Println(blogId)
}
func DeleteComment(blogId string, commentId string){
	log.Println(blogId, commentId)
}

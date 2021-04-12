package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
	"yxcblog/dbInit"
	"yxcblog/model"
	"yxcblog/model/request"
)

func CreateBlog(blogForm request.BlogForm)error{
	DB := dbInit.DB
	id := sha256.Sum256([]byte(string(rune(time.Now().Nanosecond()))+ strconv.FormatInt(rand.Int63(), 10)))
	blogId := id[:]
	result, err := DB.Exec("INSERT INTO blog(`createAt`,`updateAt`,`title`,`content`,`type`,`label`,`visitedN`,`commentN`,`blogId`,`public`) VALUES(?,?,?,?,?,?,?,?,?,?)",
		time.Now(),time.Now(),blogForm.Title,blogForm.Content,blogForm.Type,blogForm.Label,blogForm.VisitedN,blogForm.CommentN,hex.EncodeToString(blogId),blogForm.Public)
	if err!=nil{
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil{
		return err
	}
	if rows == 1{
		return nil
	}else{
		return errors.New("")
	}
}
func CreateComment(blogId string, fatherCommentId string){
	log.Println(blogId, fatherCommentId)
}

func ReadBlogById(blogId string)(gin.H, error){
	DB := dbInit.DB
	var blog model.Blog
	row := DB.QueryRow("SELECT createAt,updateAt,deleteAt, title,`content`, `type`, label, visitedN,blogId, commentN,`public` FROM blog WHERE blogId=?;", blogId)
	err := row.Scan(&blog.CreateAt, &blog.UpdateAt, &blog.DeleteAt, &blog.Title, &blog.Content, &blog.Type, &blog.Label, &blog.VisitedN, &blog.BlogId, &blog.CommentN, &blog.Public)
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
		"public": blog.Public,
	}
	return ans, nil
}
func ReadBlogByPage(pageNum int, pageSize int)([]gin.H, error){
	DB := dbInit.DB
	ans := make([]gin.H, 0)
	rows, err := DB.Query("SELECT `id`, createAt, updateAt, `title`,`content`,`type`, `label`, visitedN,blogId,commentN,`public` FROM blog ORDER BY createAt LIMIT ?,?;", (pageNum-1)*pageSize, pageNum*pageSize)
	if err != nil{
		return ans, err
	}
	for rows.Next(){
		var row model.Blog
		err := rows.Scan(&row.ID, &row.CreateAt, &row.UpdateAt, &row.Title, &row.Content, &row.Type, &row.Label, &row.VisitedN, &row.BlogId, &row.CommentN, &row.Public)
		if err!=nil{
			return ans, err
		}
		ans = append(ans, gin.H{
			"id": row.ID,
			"createAt":row.CreateAt,
			"updateAt":row.UpdateAt,
			"title": row.Title,
			"content":row.Content,
			"type": row.Type,
			"label": row.Label,
			"visitedN": row.VisitedN,
			"blogId":row.BlogId,
			"commentN": row.CommentN,
			"public": row.Public,
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
	result, err := DB.Exec("UPDATE blog SET `title`=?,`content`=?,`visitedN`=?,`commentN`=?,`type`=?,`label`=?,`public`=? WHERE `blogId`=?", form.Title, form.Content, form.VisitedN, form.CommentN, form.Type, form.Label, form.Public, blogId)
	if err != nil{
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil{
		return err
	}
	if rows == 1{
		return nil
	}else{
		return errors.New("")
	}
}

func DeleteBlogById(blogId string) error{
	DB := dbInit.DB
	result, err := DB.Exec("DELETE FROM blog WHERE blogId=?",blogId)
	if err!=nil{
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil{
		return err
	}
	if rows == 1{
		return nil
	}else{
		return errors.New("")
	}
}
func DeleteComment(blogId string, commentId string){
	log.Println(blogId, commentId)
}

package dao

import (
	"log"
	"time"
)
import "yxcblog/dbInit"

type User struct{
	ID int
	AddTime string
	Power int
	Username string
	Nickname string
	Password string
	LabelId string
	Mobile string
	Introducer int
	Email string
	LastLogin time.Time
}

func AddUser(user User)error{
	db := dbInit.DB
	result, err := db.Exec("INSERT INTO user (add_time, power, username, nickname, password, label_id, mobile, introducer, email, last_login) VALUES (?,1000,?,'',?,?,?,?,?,?)",
		user.AddTime,
		user.Username,
		user.Password,
		user.LabelId,
		user.Mobile,
		user.Introducer,
		user.Email,
		user.LastLogin,
		)
	if err!=nil{
		return err
	}
	log.Println(result)
	return nil
}
func QueryByUserId(){

}
func QueryByUserUserName(){

}
func QueryByUserEmail(){

}
func QueryByUserMobile(){

}

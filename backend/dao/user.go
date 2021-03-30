package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"yxcblog/utils"
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
	LastLogin string
}

func AddUser(user User)error{
	db := dbInit.DB
	log.Printf("%x", user.Password)
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
func QueryByUserUserName(username string, password string)(string,error){
	db := dbInit.DB
	log.Println(username, password)

	var user User

	pwd := fmt.Sprintf("%x", utils.Sha256(password))
	result := db.QueryRow("SELECT `id`,`power`,`username`,`nickname` FROM user WHERE username=? and password=?",username, pwd)

	err := result.Scan(&user.ID, &user.Power, &user.Username, &user.Nickname)
	if err!=nil{
		return "", err
	}

	ans, err := json.Marshal(map[string]interface{}{
		"id": user.ID,
		"power": user.Power,
		"username": user.Username,
		"nickname": user.Nickname,
	})
	if err != nil{
		return "", err
	}
	log.Println(ans)
	return string(ans), nil
}
func QueryByUserEmail(){

}
func QueryByUserMobile(){

}

package database

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB


func init(){
	sqlDBPool, err := sql.Open("mysql", "root:yangxuechao123@/medicalsystem?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Println(err)
		return
	}
	sqlDBPool.SetMaxOpenConns(200)
	sqlDBPool.SetMaxIdleConns(100)
	sqlDBPool.SetConnMaxLifetime(time.Hour)
	DB = sqlDBPool
}

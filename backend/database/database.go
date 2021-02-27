package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init(){
	sqlDBPool, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, IPADDRESS, PORT, DATABASE))
	if err!=nil{
		log.Println(err)
		return
	}
	sqlDBPool.SetMaxOpenConns(200)
	sqlDBPool.SetMaxIdleConns(100)
	sqlDBPool.SetConnMaxLifetime(time.Hour)
	DB = sqlDBPool
}

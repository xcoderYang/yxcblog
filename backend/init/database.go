package init

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"yxcblog/config"
)

var DB *sql.DB

func init(){
	sqlDBPool, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.MYSQL_USERNAME, config.MYSQL_PASSWORD, config.IPADDRESS, config.MYSQL_PORT, config.MYSQL_DATABASE))
	if err!=nil{
		log.Println(err)
		return
	}
	sqlDBPool.SetMaxOpenConns(200)
	sqlDBPool.SetMaxIdleConns(100)
	sqlDBPool.SetConnMaxLifetime(time.Hour)
	DB = sqlDBPool
}
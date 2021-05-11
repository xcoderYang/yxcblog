package dbInit

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	config "yxcblog/config"
)

var DB *sql.DB

func init(){
	sqlDBPool, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", config.MYSQL_USERNAME, config.MYSQL_PASSWORD, config.IPADDRESS, config.MYSQL_PORT, config.MYSQL_DATABASE))
	if err!=nil{
		log.Println(err)
		return
	}
	sqlDBPool.SetMaxOpenConns(10)
	sqlDBPool.SetMaxIdleConns(10)
	sqlDBPool.SetConnMaxLifetime(time.Minute*3)
	DB = sqlDBPool
}
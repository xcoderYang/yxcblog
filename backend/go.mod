module yxcblog

go 1.14

require (
	//router v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/tidwall/gjson v1.7.4
)

//replace router => ./router

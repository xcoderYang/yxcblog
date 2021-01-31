module yxcblog

go 1.14

require (
	//router v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

//replace router => ./router

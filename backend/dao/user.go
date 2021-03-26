package dao

import "time"

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

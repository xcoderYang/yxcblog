package model

import "time"

type Blog struct{
	ID int
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Article
}

type Article struct{
	Title string
	Content string
	WorkNum int
	Type string
	Order int
	Label string
	VisitedN int
	BlogId string
	CommentN int
}

type Comment struct{
	ID int
	AuthorId int
	Time string
	ReplyTo int
	Blog int
	BlackList bool
}

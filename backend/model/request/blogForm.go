package request

type BlogForm struct{
	BlogId string `json:"blogId"`
	Id string	`json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	VisitedN string `json:"visitedN"`
	CommentN string `json:"commentN"`
	Type string `json:"type"`
	Label string `json:"label"`
	Public string `json:"public"`
}
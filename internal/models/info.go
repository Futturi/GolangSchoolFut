package models

type Information struct {
	Author       string `json:"author"`
	About        string `json:"about"`
	CountUsers   string `json:"count_users"`
	CountLessons string `json:"count_lessons"`
}

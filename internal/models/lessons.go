package models

type Lesson struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Filling string `json:"filling" db:"filling"`
}

// type AllLessons struct {
// 	Data []Lesson `json:"lessons"`
// }

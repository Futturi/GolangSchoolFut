package models

import "errors"

type Lesson struct {
	Id       int      `json:"id" db:"id"`
	Title    string   `json:"title" db:"title"`
	Filling  string   `json:"filling" db:"filling"`
	Students []string `json:"students" db:"students"`
	Homework Homework
}

type LessonUser struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Filling  string `json:"filling" db:"filling"`
	Homework Homework
}

type LessonUserWithHome struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Filling    string `json:"filling" db:"filling"`
	HomeworkId int    `db:"homework_id"`
}

type LessonFile struct {
	FileName string
	FilePath string
}

type UpdateLesson struct {
	Title   *string `json:"title" db:"title"`
	Filling *string `json:"filling" db:"filling"`
}

func (u UpdateLesson) Validate() error {
	if u.Title == nil && u.Filling == nil {
		return errors.New("no values in update")
	}
	return nil
}

type Homework struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Descript string `json:"descript" db:"descript"`
	Mark     int    `json:"mark" db:"mark"`
}

type HomeworkUser struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Descript string `json:"descript" db:"descript"`
}

type CheckHom struct {
	Mark        int `json:"mark" db:"mark"`
	Homework_Id int `json:"homework_id" db:"homework_id"`
}

package models

import "errors"

type Lesson struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Filling  string `json:"filling" db:"filling"`
	Homework Homework
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
}

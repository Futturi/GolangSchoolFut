package models

import "errors"

type Lesson struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Filling string `json:"filling" db:"filling"`
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
		errors.New("no values in update")
	}
	return nil
}

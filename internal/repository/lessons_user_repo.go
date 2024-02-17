package repository

import (
	"fmt"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
)

type Lessons_User struct {
	db *sqlx.DB
}

func NewLessons_User(db *sqlx.DB) *Lessons_User {
	return &Lessons_User{db: db}
}

func (r *Lessons_User) GetAllLessonsuser(user_id int) ([]models.LessonUser, error) {
	var lessons []models.LessonUser

	query := fmt.Sprintf("SELECT title, filling FROM %s l INNER JOIN %s lu ON lu.lesson_id = l.id WHERE lu.user_id = $1", lessonsTable, lesson_userTable)
	err := r.db.Select(&lessons, query, user_id)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

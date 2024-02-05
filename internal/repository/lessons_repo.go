package repository

import (
	"fmt"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
)

type LessonRepo struct {
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

func (r *LessonRepo) GetAllLessonsTeacher(id int) ([]models.Lesson, error) {
	var lessons []models.Lesson
	query := fmt.Sprintf("SELECT t.id, t.title, t.filling from %s t INNER JOIN %s lt ON t.id = lt.lesson_id WHERE lt.teacher_id = $1", lessonsTable, lesson_teacher_table)
	err := r.db.Select(&lessons, query, id)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *LessonRepo) CreateLesson(userId int, mod models.Lesson) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s(title, filling) VALUES ($1, $2) RETURNING id ", lessonsTable)
	row := tx.QueryRow(query, mod.Title, mod.Filling)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	queryr := fmt.Sprintf("INSERT INTO %s(lesson_id, teacher_id) VALUES($1, $2)", lesson_teacher_table)
	_, err = tx.Query(queryr, id, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *LessonRepo) DeleteLesson(user, lesson_id int) error {
	query := fmt.Sprintf("DELETE FROM %s l USING %s tl WHERE l.id = tl.teacher_id AND l.id = $1 AND tl.teacher_id = $2", lessonsTable, lesson_teacher_table)
	_, err := r.db.Exec(query, lesson_id, user)
	if err != nil {
		return err
	}
	return nil
}

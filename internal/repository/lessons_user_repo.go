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

	query := fmt.Sprintf("SELECT id, title, filling FROM %s l INNER JOIN %s lu ON lu.lesson_id = l.id WHERE lu.user_id = $1", lessonsTable, lesson_userTable)
	err := r.db.Select(&lessons, query, user_id)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *Lessons_User) GetLessonUser(user_id, lesson_id int) (models.LessonUser, error) {
	var lesson models.LessonUserWithHome
	var hm models.Homework
	tx, err := r.db.Begin()
	if err != nil {
		return models.LessonUser{}, err
	}
	query1 := fmt.Sprintf("SELECT id, title, filling, homework_id FROM %s l WHERE l.id = $1", lessonsTable)
	row := r.db.QueryRow(query1, lesson_id)
	err = row.Scan(&hm.Id, &lesson.Title, &lesson.Filling, &lesson.HomeworkId)
	if err != nil {
		tx.Rollback()
		return models.LessonUser{}, err
	}
	query2 := fmt.Sprintf("SELECT id, title, descript, mark FROM %s h INNER JOIN %s hu ON hu.homework_id = h.id WHERE hu.user_id = $1 AND hu.homework_id = $2", homeworkTable, homeworks_userTable)
	row = r.db.QueryRow(query2, user_id, lesson.HomeworkId)
	err = row.Scan(&hm.Id, &hm.Title, &hm.Descript, &hm.Mark)
	if err != nil {
		tx.Rollback()
		return models.LessonUser{}, err
	}
	result := models.LessonUser{
		Id:       lesson.Id,
		Title:    lesson.Title,
		Filling:  lesson.Filling,
		Homework: hm,
	}
	return result, nil
}

func (r *Lessons_User) SolveHomework(user_id, lesson_id int, hw models.HomeworkUser) error {
	var hid int
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query2 := fmt.Sprintf("INSERT INTO %s(title, descript) VALUES($1, $2) RETURNING id", homeworkTable)
	row := tx.QueryRow(query2, hw.Title, hw.Descript)
	if err = row.Scan(&hid); err != nil {
		tx.Rollback()
		return err
	}
	query3 := fmt.Sprintf("INSERT INTO %s(homework_id, lesson_id) VALUES($1, $2)", lesson_homeworksTable)
	fmt.Println(hid, lesson_id)
	_, err = tx.Exec(query3, hid, lesson_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

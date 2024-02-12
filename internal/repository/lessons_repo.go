package repository

import (
	"fmt"
	"strings"

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

func (r *LessonRepo) GetLesson(id, lesson_id int) (models.Lesson, error) {
	var result models.Lesson
	query := fmt.Sprintf(`SELECT l.id, l.title, l.filling FROM %s
	 l INNER JOIN %s tl ON l.id = tl.lesson_id WHERE l.id = tl.lesson_id AND tl.teacher_id = $1
	 AND l.id = $2`, lessonsTable, lesson_teacher_table)

	row := r.db.QueryRow(query, id, lesson_id)
	if err := row.Scan(&result.Id, &result.Title, &result.Filling); err != nil {
		return models.Lesson{}, err
	}
	return result, nil
}

func (r *LessonRepo) UpdateLesson(id, lesson_id int, fil models.UpdateLesson) (models.UpdateLesson, error) {
	args := make([]interface{}, 0)
	setVal := make([]string, 0)
	argid := 1
	if fil.Title != nil {
		setVal = append(setVal, fmt.Sprintf("title=$%d", argid))
		args = append(args, *fil.Title)
		argid++
	}
	if fil.Filling != nil {
		setVal = append(setVal, fmt.Sprintf("filling=$%d", argid))
		args = append(args, *(fil.Filling))
		argid++
	}
	setQuery := strings.Join(setVal, ",")
	query := fmt.Sprintf("UPDATE %s l SET %s FROM %s tl WHERE l.id = tl.lesson_id AND l.id = $%d AND tl.teacher_id = $%d", lessonsTable, setQuery, lesson_teacher_table, argid, argid+1)
	args = append(args, lesson_id, id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return models.UpdateLesson{}, err
	}
	return fil, nil
}

func (r *LessonRepo) CreateHomework(homework models.Homework, lesson_id int) (string, error) {
	var idhom int
	tx, err := r.db.Begin()
	if err != nil {
		return "bad", err
	}

	query1 := fmt.Sprintf("INSERT INTO %s(title, descript, lesson_id) VALUES ($1, $2, $3) RETURNING id", homeworkTable)

	row := tx.QueryRow(query1, homework.Title, homework.Descript, lesson_id)
	err = row.Scan(&idhom)
	if err != nil {
		tx.Rollback()
		return "bad", err
	}

	query2 := fmt.Sprintf("UPDATE %s SET homework_id = $1 WHERE id = $2", lessonsTable)
	_, err = tx.Exec(query2, idhom, lesson_id)
	if err != nil {
		tx.Rollback()
		return "bad", err
	}
	return "good", tx.Commit()
}

func (r *LessonRepo) PutFile(name string, lesson_id int) error {
	query := fmt.Sprintf("UPDATE %s SET filename = $1 WHERE id = $2", lessonsTable)
	_, err := r.db.Exec(query, name, lesson_id)
	if err != nil {
		return err
	}
	return nil
}

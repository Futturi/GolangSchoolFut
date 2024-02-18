package repository

import (
	"errors"
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
	var user int
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
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	tx2, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	for _, stud := range mod.Students {
		queryrr := fmt.Sprintf("SELECT id FROM %s WHERE username = $1", studentTable)
		row := tx2.QueryRow(queryrr, stud)
		if err = row.Scan(&user); err != nil {
			tx2.Rollback()
			return 0, errors.New("you entered wrong username users")
		}
		queryr4 := fmt.Sprintf("INSERT INTO %s(lesson_id, user_id) VALUES($1, $2)", lesson_userTable)
		_, err = tx2.Exec(queryr4, id, user)
		if err != nil {
			tx2.Rollback()
			return 0, err
		}
	}
	return id, tx2.Commit()
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

// func (r *LessonRepo) CreateHomework(homework models.Homework, lesson_id int) (string, error) {
// 	var idhom int
// 	var iduser int
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return "bad", err
// 	}

// 	query1 := fmt.Sprintf("INSERT INTO %s(title, descript, lesson_id) VALUES ($1, $2, $3) RETURNING id", homeworkTable)

// 	row := tx.QueryRow(query1, homework.Title, homework.Descript, lesson_id)
// 	err = row.Scan(&idhom)
// 	if err != nil {
// 		tx.Rollback()
// 		return "bad", err
// 	}

// 	query2 := fmt.Sprintf("UPDATE %s SET homework_id = $1 WHERE id = $2", lessonsTable)
// 	_, err = tx.Exec(query2, idhom, lesson_id)
// 	if err != nil {
// 		tx.Rollback()
// 		return "bad", err
// 	}
// 	if err = tx.Commit(); err != nil {
// 		return "bad", err
// 	}
// 	tx2, err := r.db.Begin()
// 	if err != nil {
// 		return "bad", err
// 	}
// 	query3 := fmt.Sprintf("SELECT user_id FROM %s WHERE lesson_id = $1", lesson_userTable)
// 	row3 := tx2.QueryRow(query3, lesson_id)
// 	if err = row3.Scan(&iduser); err != nil {
// 		tx2.Rollback()
// 		return "bad", err
// 	}
// 	query4 := fmt.Sprintf("INSERT INTO %s(homework_id, user_id) VALUES($1, $2)", homeworks_userTable)
// 	_, err = tx2.Exec(query4, idhom, iduser)
// 	if err != nil {
// 		tx2.Rollback()
// 		return "bad", err
// 	}
// 	return "good", tx2.Commit()
// }

func (r *LessonRepo) PutFile(name string, lesson_id int) error {
	query := fmt.Sprintf("UPDATE %s SET filename = $1 WHERE id = $2", lessonsTable)
	_, err := r.db.Exec(query, name, lesson_id)
	if err != nil {
		return err
	}
	return nil
}
func (r *LessonRepo) CheckHomework(teacher_id, lesson_id int, status models.CheckHom) error {
	query := fmt.Sprintf("UPDATE %s SET mark = $1 FROM %s hl WHERE hl.lesson_id = $2 AND hl.homework_id = $3", homeworkTable, lesson_homeworksTable)
	_, err := r.db.Exec(query, status.Mark, lesson_id, status.Homework_Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *LessonRepo) GetHomework(lesson_id int) ([]models.Homework, error) {
	var hm []models.Homework
	query := fmt.Sprintf("SELECT id, title, descript, mark FROM %s h INNER JOIN %s hl ON hl.lesson_id = $1", homeworkTable, lesson_homeworksTable)
	err := r.db.Select(&hm, query, lesson_id)
	if err != nil {
		return []models.Homework{}, err
	}

	return hm, nil
}

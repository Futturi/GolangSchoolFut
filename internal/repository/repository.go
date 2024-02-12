package repository

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	teachersTable        = "teacher"
	lessonsTable         = "lesson"
	lesson_teacher_table = "lesson_teacher"
	studentTable         = "student"
	lesson_userTable     = "lesson_user"
	homeworkTable        = "homeworks"
	homeworks_userTable  = "homeworks_user"
)

type Repository struct {
	Authorization
	Lessons
}

type Lessons interface {
	GetAllLessonsTeacher(id int) ([]models.Lesson, error)
	CreateLesson(userId int, mod models.Lesson) (int, error)
	DeleteLesson(user, lesson_id int) error
	GetLesson(id, lesson_id int) (models.Lesson, error)
	CreateHomework(homework models.Homework, lesson_id int) (string, error)
	UpdateLesson(id, lesson_id int, fil models.UpdateLesson) (models.UpdateLesson, error)
	PutFile(name string, lesson_id int) error
}

type Authorization interface {
	SignUp(mod models.Teacher) (int, error)
	SignIn(mod models.SignInTeacher, refresh string, timerefresh int64) (int, error)
	GetByRefresh(refresh string) (int, error)
}

func NewReposiotry(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthRepo(db), Lessons: NewLessonRepo(db)}
}

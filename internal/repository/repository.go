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
	AuthorizationUser
	LessonsUser
}

type Lessons interface {
	GetAllLessonsTeacher(id int) ([]models.Lesson, error)
	CreateLesson(userId int, mod models.Lesson) (int, error)
	DeleteLesson(user, lesson_id int) error
	GetLesson(id, lesson_id int) (models.Lesson, error)
	CreateHomework(homework models.Homework, lesson_id int) (string, error)
	UpdateLesson(id, lesson_id int, fil models.UpdateLesson) (models.UpdateLesson, error)
	PutFile(name string, lesson_id int) error
	CheckHomework(teacher_id, lesson_id, status int) error
	GetHomework(lesson_id int) (models.Homework, error)
}

type Authorization interface {
	SignUp(mod models.Teacher) (int, error)
	SignIn(mod models.SignInTeacher, refresh string, timerefresh int64) (int, error)
	GetByRefresh(refresh string) (int, error)
}

type AuthorizationUser interface {
	SignUpStudent(user models.Student) (string, error)
	SignInStudent(userlog models.SignInStudent) (int, error)
}

type LessonsUser interface {
	GetAllLessonsuser(user_id int) ([]models.LessonUser, error)
}

func NewReposiotry(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:     NewAuthRepo(db),
		Lessons:           NewLessonRepo(db),
		AuthorizationUser: NewAuthorization_User(db),
		LessonsUser:       NewLessons_User(db),
	}
}

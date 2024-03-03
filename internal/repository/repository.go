package repository

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const (
	teachersTable         = "teacher"
	lessonsTable          = "lesson"
	lesson_teacher_table  = "lesson_teacher"
	studentTable          = "student"
	lesson_userTable      = "lesson_user"
	homeworkTable         = "homeworks"
	homeworks_userTable   = "homeworks_user"
	lesson_homeworksTable = "lesson_homeworks"
)

type Repository struct {
	Authorization
	Lessons
	AuthorizationUser
	LessonsUser
	Info
}

type Lessons interface {
	GetAllLessonsTeacher(id int) ([]models.Lesson, error)
	CreateLesson(userId int, mod models.Lesson) (int, error)
	DeleteLesson(user, lesson_id int) error
	GetLesson(id, lesson_id int) (models.Lesson, error)
	UpdateLesson(id, lesson_id int, fil models.UpdateLesson) (models.UpdateLesson, error)
	PutFile(name string, lesson_id int) error
	CheckHomework(teacher_id, lesson_id int, status models.CheckHom) error
	GetHomework(lesson_id int) ([]models.Homework, error)
	DecrementHealth(lesson_id int, status models.CheckHom) error
}

type Authorization interface {
	SignUp(mod models.Teacher) (int, error)
	SignIn(mod models.SignInTeacher, refresh string, timerefresh int64) (int, error)
	GetByRefresh(refresh string) (int, error)
}

type AuthorizationUser interface {
	SignUpStudent(user models.Student) (string, error)
	SignInStudent(userlog models.SignInStudent, refresh string, exp int64) (int, error)
	CheckHealth(user_id int) int
	GetIdByRefresh(refresh models.Refresh) (int, int64, error)
	CheckToken(token string) error
	CheckVer(userlog models.SignInStudent) (bool, error)
}

type LessonsUser interface {
	GetAllLessonsuser(user_id int) ([]models.LessonUser, error)
	GetLessonUser(user_id, lesson_id int) (models.LessonUser, error)
	SolveHomework(user_id, lesson_id int, hw models.HomeworkUser) error
}
type Info interface {
	Info() (models.Information, error)
}

func NewReposiotry(db *sqlx.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Authorization:     NewAuthRepo(db),
		Lessons:           NewLessonRepo(db),
		AuthorizationUser: NewAuthorization_User(db),
		LessonsUser:       NewLessons_User(db),
		Info:              NewInfo_Repo(db, rdb),
	}
}

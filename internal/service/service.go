package service

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
)

type Service struct {
	Authorization
	Lessons
	AuthorizationUser
	LessonsUser
	Infos
}
type Sender struct {
	EmailSenderName     string
	EmailSenderAddress  string
	EmailSenderPassword string
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
}

type Authorization interface {
	SignUp(mod models.Teacher) (int, error)
	SignIn(mod models.SignInTeacher) (string, string, error)
	RefreshToken(refresh string) (string, error)
	ParseToken(header string) (int, error)
}

type AuthorizationUser interface {
	SignUpStudent(user models.Student, cfg Sender) (string, error)
	SignInStudent(userlog models.SignInStudent) (string, string, error)
	ParseTokenUser(header string) (int, error)
	CheckHealth(user_id int) int
	RefreshUser(refresh models.Refresh) (string, error)
	CheckToken(token string) error
}

type LessonsUser interface {
	GetAllLessonsuser(user_id int) ([]models.LessonUser, error)
	GetLessonUser(user_id, lesson_id int) (models.LessonUser, error)
	SolveHomework(user_id, lesson_id int, hw models.HomeworkUser) error
}

type Infos interface {
	Info() (models.Information, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repo.Authorization),
		Lessons:           NewLessonsService(repo.Lessons),
		AuthorizationUser: NewAuthServiceUser(repo.AuthorizationUser),
		LessonsUser:       NewLesson_User(repo.LessonsUser),
		Infos:             NewInfoSer(repo.Info),
	}
}

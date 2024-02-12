package service

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
)

type Service struct {
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
	SignIn(mod models.SignInTeacher) (string, string, error)
	RefreshToken(refresh string) (string, error)
	ParseToken(header string) (int, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repo.Authorization), Lessons: NewLessonsService(repo.Lessons)}
}

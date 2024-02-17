package service

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
)

type Lesson_User struct {
	repo repository.LessonsUser
}

func NewLesson_User(repo repository.LessonsUser) *Lesson_User {
	return &Lesson_User{repo: repo}
}

func (s *Lesson_User) GetAllLessonsuser(user_id int) ([]models.LessonUser, error) {
	return s.repo.GetAllLessonsuser(user_id)
}

package service

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
)

type LessonsService struct {
	repo repository.Lessons
}

func NewLessonsService(repo repository.Lessons) *LessonsService {
	return &LessonsService{repo: repo}
}

func (s *LessonsService) GetAllLessonsTeacher(id int) ([]models.Lesson, error) {
	return s.repo.GetAllLessonsTeacher(id)
}

func (s *LessonsService) CreateLesson(userId int, mod models.Lesson) (int, error) {
	return s.repo.CreateLesson(userId, mod)
}

func (s *LessonsService) DeleteLesson(user, lesson_id int) error {
	return s.repo.DeleteLesson(user, lesson_id)
}

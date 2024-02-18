package service

import (
	"errors"

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

func (s *LessonsService) GetLesson(id, lesson_id int) (models.Lesson, error) {
	return s.repo.GetLesson(id, lesson_id)
}
func (s *LessonsService) UpdateLesson(id, lesson_id int, fil models.UpdateLesson) (models.UpdateLesson, error) {
	if err := fil.Validate(); err != nil {
		return models.UpdateLesson{}, err
	}
	return s.repo.UpdateLesson(id, lesson_id, fil)
}

func (s *LessonsService) PutFile(name string, lesson_id int) error {
	return s.repo.PutFile(name, lesson_id)
}

func (s *LessonsService) CheckHomework(teacher_id, lesson_id int, status models.CheckHom) error {
	if status.Mark > 5 || status.Mark < 0 {
		return errors.New("your mark is > than 5 or < than 0")
	}
	return s.repo.CheckHomework(teacher_id, lesson_id, status)
}

func (s *LessonsService) GetHomework(lesson_id int) ([]models.Homework, error) {
	return s.repo.GetHomework(lesson_id)
}

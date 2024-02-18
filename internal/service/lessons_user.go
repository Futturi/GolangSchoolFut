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

func (s *Lesson_User) GetLessonUser(user_id, lesson_id int) (models.LessonUser, error) {
	return s.repo.GetLessonUser(user_id, lesson_id)
}

func (s *Lesson_User) SolveHomework(user_id, lesson_id int, hw models.HomeworkUser) error {
	return s.repo.SolveHomework(user_id, lesson_id, hw)
}

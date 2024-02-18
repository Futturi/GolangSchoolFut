package service

import (
	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
)

type InfoSer struct {
	repo repository.Info
}

func NewInfoSer(repo repository.Info) *InfoSer {
	return &InfoSer{repo: repo}
}

func (s *InfoSer) Info() (models.Information, error) {
	return s.repo.Info()
}

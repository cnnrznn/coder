package service

import (
	"github.com/cnnrznn/coder/models"
	"github.com/cnnrznn/coder/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetProblems() ([]models.Problem, error) {
	return nil, nil
}

func (s *Service) CreateProblem(statement string) error {
	return nil
}

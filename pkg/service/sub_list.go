package service

import (
	effective "github.com/effective"
	"github.com/effective/pkg/repository"
)

type SubListService struct {
	repo repository.Subscription
}

func NewSubListService(repo repository.Subscription) *SubListService {
	return &SubListService{repo: repo}
}

func (s *SubListService) Create(sub effective.Sub) (int, error) {
	return s.repo.Create(sub)
}

package service

import (
	"github.com/effective"
	"github.com/effective/pkg/repository"
)

type SubService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubService {
	return &SubService{repo: repo}
}

func (s *SubService) Create(sub effective.Sub) (string, error) {
	return s.repo.Create(sub)
}

func (s *SubService) GetSub(id, name string) (effective.Sub, error) {
	return s.repo.GetSub(id, name)
}

func (s *SubService) GetAllSubs(id string) ([]effective.Sub, error) {
	return s.repo.GetAllSubs(id)
}

func (s *SubService) DeleteSub(id, name string) error {
	return s.repo.DeleteSub(id, name)
}

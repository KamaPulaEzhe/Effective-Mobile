package service

import (
	"github.com/effective"
	"github.com/effective/pkg/repository"
)

type Subscription interface {
	Create(sub effective.Sub) (string, error)
	GetSub(id, name string) (effective.Sub, error)
	DeleteSub(id, name string) error
	GetAllSubs(id string) ([]effective.Sub, error)
}

type Service struct {
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repos.Subscription),
	}
}

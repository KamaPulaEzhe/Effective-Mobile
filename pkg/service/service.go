package service

import (
	"github.com/effective"
	"github.com/effective/pkg/repository"
)

type Subscription interface {
	Create(sub effective.Sub) (int, error)
}

type Service struct {
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

package service

import "github.com/efective/pkg/repository"

type Subscription interface {
}

type Service struct {
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

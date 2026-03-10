package service

import (
	"github.com/effective"
	"github.com/effective/pkg/repository"
	"github.com/sirupsen/logrus"
)

type SubService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubService {
	return &SubService{repo: repo}
}

func (s *SubService) Create(sub effective.Sub) (string, error) {
	logrus.Infof("service.Create: user_id=%s service=%s", sub.UserID, sub.ServiceName)
	id, err := s.repo.Create(sub)
	if err != nil {
		logrus.Errorf("service.Create: %s", err.Error())
	}
	return id, err
}

func (s *SubService) GetSub(id, name string) (effective.Sub, error) {
	logrus.Infof("service.GetSub: user_id=%s service=%s", id, name)
	sub, err := s.repo.GetSub(id, name)
	if err != nil {
		logrus.Errorf("service.GetSub: %s", err.Error())
	}
	return sub, err
}

func (s *SubService) GetAllSubs(id string) ([]effective.Sub, error) {
	logrus.Infof("service.GetAllSubs: user_id=%s", id)
	subs, err := s.repo.GetAllSubs(id)
	if err != nil {
		logrus.Errorf("service.GetAllSubs: %s", err.Error())
	}
	return subs, err
}

func (s *SubService) DeleteSub(id, name string) error {
	logrus.Infof("service.DeleteSub: user_id=%s service=%s", id, name)
	err := s.repo.DeleteSub(id, name)
	if err != nil {
		logrus.Errorf("service.DeleteSub: %s", err.Error())
	}
	return err
}

func (s *SubService) UpdateSub(subID string, input effective.UpdateSubInput) error {
	logrus.Infof("service.UpdateSub: id=%s", subID)
	err := s.repo.UpdateSub(subID, input)
	if err != nil {
		logrus.Errorf("service.UpdateSub: %s", err.Error())
	}
	return err
}

func (s *SubService) GetTotalCost(filter effective.CostFilter) (int, error) {
	logrus.Infof("service.GetTotalCost: user_id=%s period=%s/%s", filter.UserID, filter.StartDate, filter.EndDate)
	total, err := s.repo.GetTotalCost(filter)
	if err != nil {
		logrus.Errorf("service.GetTotalCost: %s", err.Error())
	}
	return total, err
}

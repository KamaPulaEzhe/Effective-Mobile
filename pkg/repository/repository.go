package repository

import (
	"github.com/effective"
	"github.com/jmoiron/sqlx"
)

type Subscription interface {
	Create(sub effective.Sub) (string, error)
	GetSub(id, name string) (effective.Sub, error)
	DeleteSub(id, name string) error
	GetAllSubs(id string) ([]effective.Sub, error)
	UpdateSub(subID string, input effective.UpdateSubInput) error
	GetTotalCost(filter effective.CostFilter) (int, error)
}

type Repository struct {
	Subscription
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Subscription: NewSubscriptionPostgres(db),
	}
}

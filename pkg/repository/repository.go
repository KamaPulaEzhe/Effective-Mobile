package repository

import (
	"github.com/effective"
	"github.com/jmoiron/sqlx"
)

type Subscription interface {
	Create(sub effective.Sub) (int, error)
}

type Repository struct {
	Subscription
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

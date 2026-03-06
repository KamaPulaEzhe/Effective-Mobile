package repository

import "github.com/jmoiron/sqlx"

type Subscription interface {
}

type Repository struct {
	Subscription
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

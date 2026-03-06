package repository

import (
	"time"

	"github.com/effective"
	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

type SubListPostgres struct {
	db *sqlx.DB
}

func NewSubListPostgres(db *sqlx.DB) *SubListPostgres {
	return &SubListPostgres{db: db}
}

func (r *SubListPostgres) Create(sub effective.Sub) (string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return "", err
	}

	var serviceID int
	checkServiceQuery := `SELECT id FROM services WHERE name = $1`
	err = tx.Get(&serviceID, checkServiceQuery, sub.ServiceName)

	if err == pgx.ErrNoRows {
		createServiceQuery := `INSERT INTO services (name) VALUES ($1) RETURNING id`
		err = tx.Get(&serviceID, createServiceQuery, sub.ServiceName)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	var subID string
	createSubQuery := `INSERT INTO subscriptions (service_id, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	startDate, err := time.Parse("01-2006", sub.StartDate)
	if err != nil {
		return "", err
	}

	var endDate *time.Time
	if sub.EndDate != "" {
		parsedEndDate, err := time.Parse("01-2006", sub.EndDate)
		if err != nil {
			return "", err
		}
		endDate = &parsedEndDate
	}

	err = tx.Get(&subID, createSubQuery,
		serviceID,
		sub.Price,
		sub.UserID,
		startDate,
		endDate,
	)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return subID, nil

}

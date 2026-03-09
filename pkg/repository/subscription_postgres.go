package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/effective"
	"github.com/jmoiron/sqlx"
)

type SubscriptionPostgres struct {
	db *sqlx.DB
}

func NewSubscriptionPostgres(db *sqlx.DB) *SubscriptionPostgres {
	return &SubscriptionPostgres{db: db}
}

func (r *SubscriptionPostgres) Create(sub effective.Sub) (string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return "", err
	}

	var serviceID int
	checkServiceQuery := `
		SELECT 
			id 
		FROM 
			services 
		WHERE 
			name = $1`

	err = tx.Get(&serviceID, checkServiceQuery, sub.ServiceName)

	if err == sql.ErrNoRows {
		createServiceQuery := `
			INSERT INTO 
				services (name) 
			VALUES 
				($1) 
			RETURNING id`

		err = tx.Get(&serviceID, createServiceQuery, sub.ServiceName)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	var subID string
	createSubQuery := `
		INSERT INTO
				subscriptions 
				(service_id, 
				price, 
				user_id, 
				start_date,
				end_date) 
			VALUES ($1, $2, $3, $4, $5) 
			RETURNING id`

	startDate, err := time.Parse("01-2006", sub.StartDate)
	if err != nil {
		return "", err
	}

	var endDate *time.Time
	if sub.EndDate != nil {
		parsedEndDate, err := time.Parse("01-2006", *sub.EndDate)
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

func (r *SubscriptionPostgres) GetSub(id, name string) (effective.Sub, error) {
	var sub effective.Sub
	queryGetSub := `
        SELECT 
			s.id,
            sv.name as service_name,
            s.price,
            s.user_id,
            to_char(s.start_date, 'MM-YYYY') as start_date,
            to_char(s.end_date, 'MM-YYYY') as end_date
        FROM subscriptions s
        JOIN services sv ON s.service_id = sv.id
        WHERE s.user_id = $1 AND sv.name = $2
    `

	err := r.db.Get(&sub, queryGetSub, id, name)
	return sub, err
}

func (r *SubscriptionPostgres) GetAllSubs(id string) ([]effective.Sub, error) {
	var subs []effective.Sub
	queryGetAllSubs := `
        SELECT 
			s.id,
            sv.name as service_name,
            s.price,
            s.user_id,
            to_char(s.start_date, 'MM-YYYY') as start_date,
            to_char(s.end_date, 'MM-YYYY') as end_date
        FROM subscriptions s
        INNER JOIN services sv ON s.service_id = sv.id
        WHERE s.user_id = $1
    `

	err := r.db.Select(&subs, queryGetAllSubs, id)
	return subs, err
}

func (r *SubscriptionPostgres) DeleteSub(id, name string) error {
	queryGetSub := `
		DELETE FROM 
			subscriptions s 
		USING 
			services sv 
		WHERE 
			s.service_id = sv.id 
			AND s.user_id = $1 
			AND sv.name = $2
    `

	result, err := r.db.Exec(queryGetSub, id, name)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SubscriptionPostgres) UpdateSub(subID string, input effective.UpdateSubInput) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	setClauses := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.ServiceName != nil {
		var serviceID int
		err := tx.Get(&serviceID, "SELECT id FROM services WHERE name = $1", *input.ServiceName)
		if err == sql.ErrNoRows {
			err = tx.Get(&serviceID, "INSERT INTO services (name) VALUES ($1) RETURNING id", *input.ServiceName)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
		setClauses = append(setClauses, fmt.Sprintf("service_id = $%d", argID))
		args = append(args, serviceID)
		argID++
	}

	if input.Price != nil {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", argID))
		args = append(args, *input.Price)
		argID++
	}

	if input.StartDate != nil {
		startDate, _ := time.Parse("01-2006", *input.StartDate)
		setClauses = append(setClauses, fmt.Sprintf("start_date = $%d", argID))
		args = append(args, startDate)
		argID++
	}

	if input.EndDate != nil {
		endDate, _ := time.Parse("01-2006", *input.EndDate)
		setClauses = append(setClauses, fmt.Sprintf("end_date = $%d", argID))
		args = append(args, endDate)
		argID++
	}

	query := fmt.Sprintf("UPDATE subscriptions SET %s WHERE id = $%d",
		strings.Join(setClauses, ", "), argID)
	args = append(args, subID)

	result, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return tx.Commit()
}

func (r *SubscriptionPostgres) GetTotalCost(filter effective.CostFilter) (int, error) {
	args := []interface{}{}
	argID := 1

	query := `
        SELECT COALESCE(SUM(s.price), 0)
        FROM subscriptions s
        JOIN services sv ON s.service_id = sv.id
        WHERE s.user_id = $1
        AND s.start_date <= $2
        AND (s.end_date IS NULL OR s.end_date >= $3)`

	startDate, err := time.Parse("01-2006", filter.StartDate)
	if err != nil {
		return 0, errors.New("invalid start_date format, expected MM-YYYY")
	}
	endDate, err := time.Parse("01-2006", filter.EndDate)
	if err != nil {
		return 0, errors.New("invalid end_date format, expected MM-YYYY")
	}

	args = append(args, filter.UserID, endDate, startDate)
	argID = 4

	if filter.ServiceName != nil {
		query += fmt.Sprintf(" AND sv.name = $%d", argID)
		args = append(args, *filter.ServiceName)
	}

	var total int
	err = r.db.Get(&total, query, args...)
	return total, err
}

package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/effective"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SubscriptionPostgres struct {
	db *sqlx.DB
}

func NewSubscriptionPostgres(db *sqlx.DB) *SubscriptionPostgres {
	return &SubscriptionPostgres{db: db}
}

func (r *SubscriptionPostgres) Create(sub effective.Sub) (string, error) {
	logrus.Infof("repo.Create: starting for user_id=%s service=%s", sub.UserID, sub.ServiceName)

	tx, err := r.db.Beginx()
	if err != nil {
		logrus.Errorf("repo.Create: failed to begin transaction: %s", err.Error())
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
		logrus.Infof("repo.Create: service=%s not found, creating new", sub.ServiceName)
		createServiceQuery := `
			INSERT INTO 
				services (name) 
			VALUES 
				($1) 
			RETURNING id`

		err = tx.Get(&serviceID, createServiceQuery, sub.ServiceName)
		if err != nil {
			logrus.Errorf("repo.Create: failed to create service=%s: %s", sub.ServiceName, err.Error())
			tx.Rollback()
			return "", err
		}
		logrus.Infof("repo.Create: service=%s created with id=%d", sub.ServiceName, serviceID)
	} else if err != nil {
		logrus.Errorf("repo.Create: failed to check service=%s: %s", sub.ServiceName, err.Error())
		tx.Rollback()
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
		logrus.Errorf("repo.Create: invalid start_date=%s: %s", sub.StartDate, err.Error())
		tx.Rollback()
		return "", err
	}

	var endDate *time.Time
	if sub.EndDate != nil {
		parsedEndDate, err := time.Parse("01-2006", *sub.EndDate)
		if err != nil {
			logrus.Errorf("repo.Create: invalid end_date=%s: %s", *sub.EndDate, err.Error())
			tx.Rollback()
			return "", err
		}
		endDate = &parsedEndDate
	}

	err = tx.Get(&subID, createSubQuery, serviceID, sub.Price, sub.UserID, startDate, endDate)
	if err != nil {
		logrus.Errorf("repo.Create: failed to insert subscription: %s", err.Error())
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("repo.Create: failed to commit transaction: %s", err.Error())
		return "", err
	}

	logrus.Infof("repo.Create: subscription created id=%s user_id=%s service=%s", subID, sub.UserID, sub.ServiceName)
	return subID, nil
}

func (r *SubscriptionPostgres) GetSub(id, name string) (effective.Sub, error) {
	logrus.Infof("repo.GetSub: user_id=%s service=%s", id, name)

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
	if err != nil {
		logrus.Errorf("repo.GetSub: user_id=%s service=%s: %s", id, name, err.Error())
		return sub, err
	}

	logrus.Infof("repo.GetSub: found subscription id=%s", sub.ID)
	return sub, nil
}

func (r *SubscriptionPostgres) GetAllSubs(id string) ([]effective.Sub, error) {
	logrus.Infof("repo.GetAllSubs: user_id=%s", id)

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
	if err != nil {
		logrus.Errorf("repo.GetAllSubs: user_id=%s: %s", id, err.Error())
		return subs, err
	}

	logrus.Infof("repo.GetAllSubs: found %d subscriptions for user_id=%s", len(subs), id)
	return subs, nil
}

func (r *SubscriptionPostgres) DeleteSub(id, name string) error {
	logrus.Infof("repo.DeleteSub: user_id=%s service=%s", id, name)

	query := `
		DELETE FROM 
			subscriptions s 
		USING 
			services sv 
		WHERE 
			s.service_id = sv.id 
			AND s.user_id = $1 
			AND sv.name = $2
    `

	result, err := r.db.Exec(query, id, name)
	if err != nil {
		logrus.Errorf("repo.DeleteSub: user_id=%s service=%s: %s", id, name, err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("repo.DeleteSub: failed to get rows affected: %s", err.Error())
		return err
	}

	if rows == 0 {
		logrus.Warnf("repo.DeleteSub: not found user_id=%s service=%s", id, name)
		return sql.ErrNoRows
	}

	logrus.Infof("repo.DeleteSub: deleted user_id=%s service=%s", id, name)
	return nil
}

func (r *SubscriptionPostgres) UpdateSub(subID string, input effective.UpdateSubInput) error {
	logrus.Infof("repo.UpdateSub: id=%s", subID)

	tx, err := r.db.Beginx()
	if err != nil {
		logrus.Errorf("repo.UpdateSub: failed to begin transaction: %s", err.Error())
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
			logrus.Infof("repo.UpdateSub: service=%s not found, creating new", *input.ServiceName)
			err = tx.Get(&serviceID, "INSERT INTO services (name) VALUES ($1) RETURNING id", *input.ServiceName)
			if err != nil {
				logrus.Errorf("repo.UpdateSub: failed to create service=%s: %s", *input.ServiceName, err.Error())
				return err
			}
		} else if err != nil {
			logrus.Errorf("repo.UpdateSub: failed to check service=%s: %s", *input.ServiceName, err.Error())
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
		logrus.Errorf("repo.UpdateSub: failed to execute update id=%s: %s", subID, err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		logrus.Errorf("repo.UpdateSub: failed to get rows affected: %s", err.Error())
		return err
	}
	if rows == 0 {
		logrus.Warnf("repo.UpdateSub: not found id=%s", subID)
		return sql.ErrNoRows
	}

	if err = tx.Commit(); err != nil {
		logrus.Errorf("repo.UpdateSub: failed to commit transaction: %s", err.Error())
		return err
	}

	logrus.Infof("repo.UpdateSub: updated id=%s", subID)
	return nil
}

func (r *SubscriptionPostgres) GetTotalCost(filter effective.CostFilter) (int, error) {
	logrus.Infof("repo.GetTotalCost: user_id=%s period=%s/%s service=%v", filter.UserID, filter.StartDate, filter.EndDate, filter.ServiceName)

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
		logrus.Errorf("repo.GetTotalCost: invalid start_date=%s: %s", filter.StartDate, err.Error())
		return 0, errors.New("invalid start_date format, expected MM-YYYY")
	}
	endDate, err := time.Parse("01-2006", filter.EndDate)
	if err != nil {
		logrus.Errorf("repo.GetTotalCost: invalid end_date=%s: %s", filter.EndDate, err.Error())
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
	if err != nil {
		logrus.Errorf("repo.GetTotalCost: user_id=%s: %s", filter.UserID, err.Error())
		return 0, err
	}

	logrus.Infof("repo.GetTotalCost: total=%d for user_id=%s", total, filter.UserID)
	return total, nil
}

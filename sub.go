package effective

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Sub struct {
	ServiceName string    `json:"service_name" db:"service_name" binding:"required"`
	Price       int       `json:"price" db:"price" binding:"required"`
	UserID      uuid.UUID `json:"user_id" db:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" db:"start_date" binding:"required"`
	// StartDate string `json:"start_date"`
	EndDate *string `json:"end_date" db:"end_date"`
	ID      string  `json:"id" db:"id"`
}

type UpdateSubInput struct {
	ServiceName *string `json:"service_name"`
	Price       *int    `json:"price"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
}

type CostFilter struct {
	UserID      string  `json:"user_id"`
	ServiceName *string `json:"service_name"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
}

type TotalCostOutput struct {
	Total int `json:"total"`
}

func (i UpdateSubInput) Validate() error {
	if i.ServiceName == nil && i.Price == nil && i.StartDate == nil && i.EndDate == nil {
		return errors.New("update structure has no values")
	}

	if i.Price != nil && *i.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if i.StartDate != nil {
		if _, err := time.Parse("01-2006", *i.StartDate); err != nil {
			return errors.New("invalid start_date format, expected MM-YYYY")
		}
	}

	if i.EndDate != nil {
		if _, err := time.Parse("01-2006", *i.EndDate); err != nil {
			return errors.New("invalid end_date format, expected MM-YYYY")
		}
	}

	return nil
}

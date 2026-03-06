package effective

import "github.com/google/uuid"

type Sub struct {
	ServiceName string    `json:"service_name" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	// StartDate   string    `json:"start_date" binding:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

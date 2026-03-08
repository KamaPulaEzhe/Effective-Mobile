package effective

import "github.com/google/uuid"

type Sub struct {
	ServiceName string    `json:"service_name" db:"service_name" binding:"required"`
	Price       int       `json:"price" db:"price" binding:"required"`
	UserID      uuid.UUID `json:"user_id" db:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" db:"start_date" binding:"required"`
	// StartDate string `json:"start_date"`
	EndDate *string `json:"end_date" db:"end_date"`
}

// type UpdateSubInput struct {
//     ServiceName *string    `json:"service_name"`
//     Price       *int       `json:"price"`
//     StartDate   *string    `json:"start_date"`
//     EndDate     *string    `json:"end_date"`
// }

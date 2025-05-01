package company

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Password    string
	Phone       string
	Description *string
	AccountType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

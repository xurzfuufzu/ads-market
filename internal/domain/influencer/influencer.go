package influencer

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
	Platforms   []string
	Category    []string
	City        *string
	AccountType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

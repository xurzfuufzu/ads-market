package influencer

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Phone       string     `json:"phone"`
	Platforms   []string   `json:"platforms"`
	Category    []string   `json:"category"`
	City        *string    `json:"city,omitempty"`
	AccountType string     `json:"account_type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

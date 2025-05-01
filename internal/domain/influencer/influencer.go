package influencer

import (
	"Ads-marketplace/internal/domain"
	"github.com/google/uuid"
)

type Entity struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Password    string
	Phone       string
	Platforms   []domain.Platform
	AccountType string
	CreatedAt   string
	UpdatedAt   string
}

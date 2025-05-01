package ad

import (
	"Ads-marketplace/internal/domain"
	"database/sql"
	"time"
)

const (
	AdStatusAvailable  = "available"
	AdStatusInProgress = "in_progress"
	AdStatusDone       = "done"
)

var validStatuses = map[string]struct{}{
	AdStatusAvailable:  {},
	AdStatusInProgress: {},
	AdStatusDone:       {},
}

type Entity struct {
	ID            string
	Title         string
	CompanyName   string
	InfluencerId  sql.NullString
	Description   sql.NullString
	Price         uint32
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Platform      []domain.Platform
	Category      sql.NullString
	TargetCountry sql.NullString
}

func IsValidStatus(status string) bool {
	_, ok := validStatuses[status]
	return ok
}

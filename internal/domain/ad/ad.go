package ad

import (
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
	ID             string
	Title          string  `json:"title"`
	CompanyName    string  `json:"company_name"`
	Description    *string `json:"description"`
	PriceFrom      uint32  `json:"priceFrom"`
	PriceTo        uint32  `json:"priceTo"`
	Status         string  `json:"status"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Platforms      []string `json:"platforms"`
	Category       *string  `json:"category"`
	City           *string  `json:"target_city"`
	ResponsesCount int      `json:"responses_count"`
}

func IsValidStatus(status string) bool {
	_, ok := validStatuses[status]
	return ok
}

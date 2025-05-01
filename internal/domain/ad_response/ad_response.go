package ad_response

import "time"

type Entity struct {
	ID           string
	AdID         string
	InfluencerID string
	Message      *string
	Status       string
	CreatedAt    time.Time
}

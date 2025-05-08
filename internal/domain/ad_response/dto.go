package ad_response

import "github.com/google/uuid"

type CreateRequest struct {
	AdID         string  `json:"ad_id"`
	InfluencerID string  `json:"influencer_id"`
	Message      *string `json:"message"`
}

type AdResponseDTO struct {
	AdsID       string `json:"ads_id"`
	Title       string `json:"title"`
	CompanyName string `json:"company_name"`
	ID          string `json:"id"`
	Status      string `json:"status"`
}

type UpdateAdStatusDTO struct {
	ID     uuid.UUID `json:"ad_response_id"`
	Status string    `json:"status"`
}

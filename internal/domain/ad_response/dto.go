package ad_response

type CreateRequest struct {
	AdID         string  `json:"ad_id"`
	InfluencerID string  `json:"influencer_id"`
	Message      *string `json:"message"`
}

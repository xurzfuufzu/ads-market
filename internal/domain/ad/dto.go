package ad

import "Ads-marketplace/internal/domain"

type CreateRequest struct {
	CompanyName string  `json:"company_name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Platforms   []domain.Platform
}

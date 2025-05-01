package ad

type CreateRequest struct {
	Title       string   `json:"title"`
	CompanyName string   `json:"company_name"`
	Description *string  `json:"description"`
	PriceFrom   uint32   `json:"priceFrom"`
	PriceTo     uint32   `json:"priceTo"`
	Platforms   []string `json:"platforms"`
	Category    *string  `json:"category"`
	City        *string  `json:"target_city"`
}

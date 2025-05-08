package influencer

import (
	"errors"
)

var (
	ErrorNotFound            = errors.New("error not found")
	ErrorEmailConflict       = errors.New("email already exists")
	ErrorInvalidName         = errors.New("name is invalid or empty")
	ErrorInvalidEmail        = errors.New("email format is invalid")
	ErrorInvalidPlatforms    = errors.New("platforms must contain at least one valid URL")
	ErrorInvalidCategory     = errors.New("category is required")
	ErrorInvalidFollowers    = errors.New("followers must be greater than zero")
	ErrorInvalidPricePerPost = errors.New("price per post must be greater than zero")
)

type RegisterRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Category    string `json:"category"`
}

type InfluencerDTO struct {
	ID        string
	Name      *string
	Platforms *[]string
	Category  *[]string
	City      *string
	Status    *string
}

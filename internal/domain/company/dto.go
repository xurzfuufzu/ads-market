package company

import "errors"

var (
	ErrorNotFound        = errors.New("error not found")
	ErrorEmailConflict   = errors.New("email already exists")
	ErrorInvalidName     = errors.New("name is invalid or empty")
	ErrorInvalidEmail    = errors.New("email format is invalid")
	ErrorInvalidPassword = errors.New("password must be at least 8 characters")
)

type RegisterRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

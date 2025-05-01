package domain

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	token string `json:"token"`
}

type Platform struct {
	Name      string `json:"name"`
	Followers uint32 `json:"followers"`
}

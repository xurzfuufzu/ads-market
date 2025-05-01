package token

import (
	"Ads-marketplace/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secretKey       []byte
	tokenTTL        time.Duration
	refreshTokenTTL time.Duration
)

func Init(cfg *config.Config) {
	secretKey = []byte(cfg.JWT.SecretKey)
	tokenTTL = cfg.JWT.TokenTTL
	refreshTokenTTL = cfg.JWT.RefreshTokenTTL
}

type TokenClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id string) (string, error) {
	claims := TokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	return claims.ID, nil
}

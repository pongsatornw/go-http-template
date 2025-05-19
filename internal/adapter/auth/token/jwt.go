package token

import (
	"context"
	"errors"
	"pongsatorn/go-http-template/internal/adapter/config"
	"pongsatorn/go-http-template/internal/core/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtManager struct {
	SecretKey []byte
	Expired   time.Duration
}

func NewJwtManager(config config.JwtConfig) *JwtManager {
	return &JwtManager{
		SecretKey: []byte(config.Secret),
		Expired:   config.Expired,
	}
}

func (s *JwtManager) Generate(ctx context.Context, userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.Expired)),
		Subject:   userID,
	})

	return jwtToken.SignedString([]byte(s.SecretKey))
}

func (s *JwtManager) Validate(ctx context.Context, tokenStr string) (string, error) {
	tok, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})

	switch {
	case errors.Is(err, domain.ErrTokenExpired):
		return "", err
	case err != nil || !tok.Valid:
		return "", domain.ErrTokenInvalid
	}

	claims, ok := tok.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", domain.ErrTokenInvalid
	}

	return claims.Subject, nil
}

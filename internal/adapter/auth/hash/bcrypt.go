package hash

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type BcryptManager struct{}

func NewBcryptManager() *BcryptManager {
	return &BcryptManager{}
}

func (m *BcryptManager) Hash(ctx context.Context, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (m *BcryptManager) Compare(ctx context.Context, hash string, plaintext string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
}

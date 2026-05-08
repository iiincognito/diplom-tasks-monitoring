package auth_service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenExpiry = 8 * time.Hour
)

type AuthService struct {
	password string
	secret   string
}

func NewAuthService() *AuthService {
	return &AuthService{
		password: os.Getenv("TODO_PASSWORD"),
		secret:   getSecret(),
	}
}

func (s *AuthService) IsAuthRequired() bool {
	return s.password != ""
}

func (s *AuthService) Authenticate(password string) (string, error) {
	if password != s.password {
		return "", fmt.Errorf("Неверный пароль")
	}

	token, err := s.generateToken()
	if err != nil {
		return "", fmt.Errorf("ошибка генерации токена: %w", err)
	}

	return token, nil
}

func (s *AuthService) ValidateToken(tokenString string) bool {
	if !s.IsAuthRequired() {
		return true
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	// Check password hash in claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	hash, ok := claims["hash"].(string)
	if !ok {
		return false
	}

	return hash == s.getPasswordHash()
}

func (s *AuthService) generateToken() (string, error) {
	claims := jwt.MapClaims{
		"hash": s.getPasswordHash(),
		"exp":  time.Now().Add(tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *AuthService) getPasswordHash() string {
	hash := sha256.Sum256([]byte(s.password))
	return hex.EncodeToString(hash[:])
}

func getSecret() string {
	secret := os.Getenv("TODO_SECRET")
	if secret == "" {
		return "default-secret-key-change-in-production"
	}
	return secret
}

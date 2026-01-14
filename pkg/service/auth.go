package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	todo "github.com/BudjakovDmitry/go_todo_app"
	"github.com/BudjakovDmitry/go_todo_app/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "strong_salt_here"
	signingKey = "sekret_key_here"
)

type tokenClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(
		username, generatePasswordHash(password),
	)
	if err != nil {
		return "", err
	}
	claims := tokenClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

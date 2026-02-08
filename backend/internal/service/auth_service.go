package service

import (
	"errors"

	"interview-copilot/backend/internal/auth"
	"interview-copilot/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Users     *repository.UserRepo
	JWTSecret string
}

func (s *AuthService) Register(email, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return s.Users.Create(email, string(hash))
}

func (s *AuthService) Login(email, password string) (string, error) {
	id, hash, err := s.Users.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}
	return auth.Generate(id, s.JWTSecret)
}

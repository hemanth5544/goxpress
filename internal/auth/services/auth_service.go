package services

import (
	"errors"

	"github.com/hemanth5544/goxpress/internal/auth/dto"

	util "github.com/hemanth5544/goxpress/internal/utils"

	"github.com/hemanth5544/goxpress/internal/auth/model"
	"github.com/hemanth5544/goxpress/internal/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.IAuthRepository
}

func NewAuthService(repo repository.IAuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(username string, email string, password string, role string) error {
	if role != "user" && role != "admin" {
		return errors.New("invalid role")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	return s.repo.CreateUser(user)

}

func (s *AuthService) Login(userLogin dto.LoginRequest) (string, error) {
	// get user if exist
	user, err := s.repo.CheckUserExist(userLogin)
	if err != nil {
		return "", errors.New("invalid email")
	}

	// hash user password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return "", errors.New("failed to compare password")
	}

	// generate token
	tokenString, err := util.GenerateToken(user.Username, user.Role)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

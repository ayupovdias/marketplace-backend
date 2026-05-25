package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(
	username,
	email,
	password string,
) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.UserRepo.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(
	email,
	password string,
) (*models.User, error) {

	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

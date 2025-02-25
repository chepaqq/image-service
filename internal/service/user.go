package service

import (
	"os"
	"strconv"
	"time"

	"github.com/chepaqq/image-service/internal/domain"
	"github.com/chepaqq/image-service/internal/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

// userService represents a service layer for user.
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates and returns a new userService object
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser create a new user
func (s *userService) CreateUser(user domain.User) (int, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(bytes)
	return s.repo.CreateUser(user)
}

// GenerateToken generates a JWT token for user
func (s *userService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUserByName(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(user.ID),
		"exp":     time.Now().Add(12 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

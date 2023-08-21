package service

import (
	"os"
	"strconv"
	"time"

	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("JWT_SECRET")

type userRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUserByName(username string) (domain.User, error)
}

// UserService represents a service layer for user.
type UserService struct {
	repo userRepository
}

// NewUserService creates and returns a new UserService object
func NewUserService(repo userRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser create a new user
func (s *UserService) CreateUser(user domain.User) (int, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(bytes)
	return s.repo.CreateUser(user)
}

// GenerateToken generates a JWT token for user
func (s *UserService) GenerateToken(username, password string) (string, error) {
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
	return token.SignedString([]byte(secretKey))
}

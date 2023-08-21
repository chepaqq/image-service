package service

import (
	"time"

	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Retrieve from env file
const (
	tokenTTL  = time.Hour * 12
	secretKey = "123213123sdfsg"
)

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
	// TODO: check if user exists already
	return s.repo.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

// GenerateToken generates a JWT token for user
func (s *UserService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUserByName(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(secretKey))
}

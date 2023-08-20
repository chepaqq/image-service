package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/golang-jwt/jwt"
)

// TODO: get signing key from env file
const signingKey = "qwertry"

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id,omitempty"`
}

type authRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUserByName(username string, passwordHash string) (domain.User, error)
}

type Service struct {
	repo authRepository
}

func NewService(repo authRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Service) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUserByName(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

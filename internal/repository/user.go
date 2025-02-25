package repository

import (
	"github.com/chepaqq/image-service/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUserByName(username string) (domain.User, error)
}

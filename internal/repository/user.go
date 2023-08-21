package repository

import (
	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/jmoiron/sqlx"
)

// UserRepository represents repository for a user entity
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates and returns a new UserRepository object
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts new user into the repository
func (r *UserRepository) CreateUser(user domain.User) (int, error) {
	var id int
	query := `INSERT INTO users(username, password_hash) VALUES ($1, $2)`
	row := r.db.QueryRow(query, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetUserByName retrieves a user from the repository
func (r *UserRepository) GetUserByName(username string) (domain.User, error) {
	var user domain.User
	query := `SELECT * from users WHERE username=$1`
	err := r.db.Get(&user, query, username)
	if err != nil {
		return user, err
	}
	return user, nil
}

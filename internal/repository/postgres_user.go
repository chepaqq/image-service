package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"github.com/chepaqq/image-service/internal/domain"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

// NewPostgresUserRepository creates and returns a new UserRepository object
func NewPostgresUserRepository(db *sqlx.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

// CreateUser inserts new user into the repository
func (r *PostgresUserRepository) CreateUser(user domain.User) (int, error) {
	var id int
	query := `INSERT INTO users(username, password_hash) VALUES ($1, $2) RETURNING id`
	result := r.db.QueryRow(query, user.Username, user.Password)
	err := result.Scan(&id)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, domain.ErrUserConflict
		}
		return 0, err
	}
	return id, nil
}

// GetUserByName retrieves a user from the repository
func (r *PostgresUserRepository) GetUserByName(username string) (domain.User, error) {
	var user domain.User
	query := `SELECT * from users WHERE username=$1`
	err := r.db.Get(&user, query, username)
	if err != nil {
		return user, err
	}
	return user, nil
}

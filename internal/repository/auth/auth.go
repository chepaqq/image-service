package auth

import (
	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user domain.User) (int, error) {
	var id int
	query := `INSERT INTO users(username, password_hash) VALUES ($1, $2)`
	row := r.db.QueryRow(query, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) GetUserByName(username string, passwordHash string) (domain.User, error) {
	var user domain.User
	query := `SELECT id from users WHERE username=$1 AND password_hash=$2`
	err := r.db.Get(&user, query, username, passwordHash)
	if err != nil {
		return user, err
	}
	return user, nil
}

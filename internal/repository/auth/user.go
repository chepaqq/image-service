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
	query := `INSERT INTO users(username, password_hash) VALUES $1, $2`
	row := r.db.QueryRow(query, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

package domain

// User represents a user entity
type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty" db:"password_hash"`
}

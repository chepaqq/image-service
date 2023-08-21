package domain

// Image repsesents an image entity
type Image struct {
	ID        int    `json:"id,omitempty"`
	UserID    int    `json:"user_id,omitempty"    db:"user_id"`
	ImagePath string `json:"image_path,omitempty" db:"image_path"`
	ImageURL  string `json:"image_url,omitempty"  db:"image_url"`
}

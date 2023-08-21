package domain

// Image repsesents an image entity
type Image struct {
	ID        int    `json:"id,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	ImagePath string `json:"image_path,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

package entity

type Post struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint64 `json:"user_id"`
}

type CreatePost struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	UserID  uint64 `json:"user_id" validate:"required"`
}

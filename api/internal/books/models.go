package books

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	Id        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	UserId    uuid.UUID  `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

package users

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id             uuid.UUID  `json:"id,omitempty"`
	UserName       string     `json:"user_name,omitempty"`
	Email          string     `json:"email,omitempty"`
	HashedPassword []byte     `json:"hashed_password"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

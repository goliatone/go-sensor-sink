package auth

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//LoginInput login payload
type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

//User model
type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username" gorm:"not null; unique_index"`
	Email     string    `json:"email" gorm:"not null; unique_index"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time
}

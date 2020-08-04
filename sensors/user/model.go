package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//User model
type User struct {
	ID        uuid.UUID `json:"id" gorm:"id" sql:"type:uuid;primary_key";default:uuid_generate_v4()`
	FirstName string    `json:"first_name" gorm:"first_name"`
	LastName  string    `json:"larst_name" gorm:"last_name"`
	CreatedAt time.Time `json:"created_at" gorm:"autocreatetime"`
	UpdatedAt time.Time `json:"update_at" gorm:"autoupdatetime"`
}

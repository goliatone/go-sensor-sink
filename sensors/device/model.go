package device

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Device holds device data
type Device struct {
	ID          uuid.UUID `json:"id" gorm:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	HardwareID  string    `json:"hardware_id" gorm:"hardware_id"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autocreatetime"`
	UpdatedAt   time.Time `json:"update_at" gorm:"autoupdatetime"`
}

//BeforeCreate hook to initialize model before creation
// func (d *Device) BeforeCreate(scope *gorm.Scope) error {
// 	id, err := uuid.NewV4().String()
// 	if err != nil {
// 		return err
// 	}

// 	return scope.SetColumn("ID", id)
// }

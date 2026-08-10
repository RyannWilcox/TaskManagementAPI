package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Resource string    `json:"resource"`
	Action   string    `json:"action"`
	Roles    []Role    `gorm:"many2many:role_permissions"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

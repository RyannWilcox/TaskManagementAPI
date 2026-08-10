package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID    `json:"id" gorm:"primaryKey"`
	Name        string       `gorm:"uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

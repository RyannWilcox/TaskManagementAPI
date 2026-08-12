package models

import "github.com/gofrs/uuid"

type Role struct {
	ID          uuid.UUID    `json:"id" gorm:"primaryKey"`
	Name        string       `gorm:"uniqueIndex;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

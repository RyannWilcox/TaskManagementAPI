package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	UserId       uuid.UUID `json:"user_id"`
	RefreshToken uuid.UUID `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

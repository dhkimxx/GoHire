package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Webhook struct {
	ID                   uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name                 string
	URL                  string
	RequiredVerification bool
	SecretKey            string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (webhook Webhook) ValidSecretKey() error {
	if webhook.RequiredVerification && len(webhook.SecretKey) == 0 {
		return errors.New("secret key is required when verification is required")
	} else {
		return nil
	}
}

func (webhook *Webhook) BeforeCreate(tx *gorm.DB) (err error) {
	err = webhook.ValidSecretKey()
	if err != nil {
		return err
	}

	webhook.ID = uuid.New()
	return nil
}

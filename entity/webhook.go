package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Webhook struct {
	ID                   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                 string
	URL                  string
	RequiredVerification bool
	SecretKey            string
}

func (webhook Webhook) ValidSecretKey() error {
	if webhook.RequiredVerification && len(webhook.SecretKey) == 0 {
		return errors.New("secret key is required when verification is required")
	} else {
		return nil
	}
}

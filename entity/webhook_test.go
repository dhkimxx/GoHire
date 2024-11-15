package entity_test

import (
	"lark-gitlab-bridge/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWebhookInitialization(t *testing.T) {
	// Given
	id := uuid.New()
	name := "Test WebHook"
	url := "https://example.com/webhook"
	requiredVerification := true
	secretKey := "supersecret"

	// When
	webhook := entity.Webhook{
		ID:                   id,
		Name:                 name,
		URL:                  url,
		RequiredVerification: requiredVerification,
		SecretKey:            secretKey,
	}

	// Then
	assert.Equal(t, id, webhook.ID)
	assert.Equal(t, name, webhook.Name)
	assert.Equal(t, url, webhook.URL)
	assert.Equal(t, requiredVerification, webhook.RequiredVerification)
	assert.Equal(t, secretKey, webhook.SecretKey)
}

func TestWebhookMustHaveSecretKeyWhenIsRequiredVerification(t *testing.T) {
	// Given
	id := uuid.New()
	name := "Test WebHook"
	url := "https://example.com/webhook"
	requiredVerification := true

	// When
	webhook := entity.Webhook{
		ID:                   id,
		Name:                 name,
		URL:                  url,
		RequiredVerification: requiredVerification,
	}

	// Then
	assert.Error(t, webhook.ValidSecretKey())
}

func TestWebhookDoesNotNeedSecretKeyWhenIsNotRequiredVerification(t *testing.T) {
	// Given
	id := uuid.New()
	name := "Test WebHook"
	url := "https://example.com/webhook"
	requiredVerification := false

	// When
	webhook := entity.Webhook{
		ID:                   id,
		Name:                 name,
		URL:                  url,
		RequiredVerification: requiredVerification,
	}

	// Then
	assert.NoError(t, webhook.ValidSecretKey())
}

func TestXxx(t *testing.T) {

}

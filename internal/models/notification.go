package models

import "github.com/google/uuid"

type InternalNotificationRequest struct {
	Recipient string
	UUID      uuid.UUID
}

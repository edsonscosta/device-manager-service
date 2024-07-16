package model

import (
	"github.com/google/uuid"
	"time"
)

type Device struct {
	DeviceID  uuid.UUID `json:"device_id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

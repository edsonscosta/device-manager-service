package repository

import (
	"device-manager-service/domain/device/model"
	"github.com/google/uuid"
)

type Device interface {
	Get(deviceID uuid.UUID) (*model.Device, error)
	GetAll(limit int, offset int, brand string) ([]*model.Device, error)
	Create(device *model.Device) error
	Update(device *model.Device) error
	UpdateStatus(deviceID uuid.UUID, status bool) error
	Delete(deviceID uuid.UUID) error
}

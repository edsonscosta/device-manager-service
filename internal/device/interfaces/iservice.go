package interfaces

import (
	"device-manager-service/internal/device/model"
	"github.com/google/uuid"
)

type IDeviceService interface {
	GetByID(deviceID uuid.UUID) (*model.Device, error)
	GetAll(limit int, offset int, brand string) ([]*model.Device, error)
	Create(device *model.Device) error
	Update(device *model.Device) error
	UpdateStatus(deviceID uuid.UUID, status bool) error
	Delete(deviceID uuid.UUID) error
}

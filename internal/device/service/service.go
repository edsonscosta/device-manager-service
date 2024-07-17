package service

import (
	"device-manager-service/internal/device/interfaces"
	"device-manager-service/internal/device/model"
	"github.com/google/uuid"
	"log"
)

type Device struct {
	deviceRepository interfaces.IDeviceRepository
	logger           *log.Logger
}

func NewDeviceService(logger *log.Logger, deviceRepository interfaces.IDeviceRepository) *Device {
	return &Device{
		deviceRepository: deviceRepository,
		logger:           logger,
	}
}

func (d Device) GetByID(deviceID uuid.UUID) (*model.Device, error) {
	return d.deviceRepository.Get(deviceID)
}

func (d Device) GetAll(limit int, offset int, brand string) ([]*model.Device, error) {
	return d.deviceRepository.GetAll(limit, offset, brand)
}

func (d Device) Create(device *model.Device) error {
	err := d.deviceRepository.Create(device)
	if err != nil {
		return err
	}
	return nil
}

func (d Device) Update(device *model.Device) error {
	err := d.deviceRepository.Update(device)
	if err != nil {
		return err
	}
	return nil
}

func (d Device) UpdateStatus(deviceID uuid.UUID, status bool) error {
	err := d.deviceRepository.UpdateStatus(deviceID, status)
	if err != nil {
		return err
	}
	return nil
}

func (d Device) Delete(deviceID uuid.UUID) error {
	err := d.deviceRepository.Delete(deviceID)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"device-manager-service/internal/device/interfaces"
	"device-manager-service/internal/device/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"os"
	"testing"
	"time"
)

var (
	logger = log.New(os.Stderr, "device-manager-service-api", 0)
)

// mockgen -destination=internal/service/mocks/mock.go -package=mocks device-manager-service/internal/device/interfaces IDeviceRepository
func TestDevice_Create(t *testing.T) {
	deviceRepo := interfaces.NewMockIDeviceRepository(t)
	deviceRepo.EXPECT().Create(mock.Anything).Return(nil)

	deviceService := NewDeviceService(logger, deviceRepo)
	deviceService.Create(
		&model.Device{
			DeviceID:  uuid.New(),
			Name:      "Name",
			Brand:     "Brand",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	assert.True(t, deviceRepo.AssertCalled(t, "Create", mock.Anything))
}

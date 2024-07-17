package handler

import (
	"device-manager-service/internal/device/interfaces"
	"device-manager-service/internal/device/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	logger = log.New(os.Stderr, "device-manager-service-api", 0)
)

func TestDevice_GetByID(t *testing.T) {
	deviceServiceMock := interfaces.NewMockIDeviceService(t)

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)

	c.Params = []gin.Param{{Key: "id", Value: "invalid_id"}}

	deviceHandler := Handler{
		deviceService: deviceServiceMock,
		logger:        logger,
	}

	deviceHandler.Get(c)
	assert.True(t, c.Writer.Status() == http.StatusBadRequest)
}

func TestDevice_GetByID_Error_Calling_Repository(t *testing.T) {
	deviceServiceMock := interfaces.NewMockIDeviceService(t)
	deviceServiceMock.EXPECT().GetByID(mock.Anything).Return(nil, fmt.Errorf("error"))

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)

	c.Params = []gin.Param{{Key: "id", Value: uuid.NewString()}}

	deviceHandler := Handler{
		deviceService: deviceServiceMock,
		logger:        logger,
	}

	deviceHandler.Get(c)
	assert.True(t, c.Writer.Status() == http.StatusInternalServerError)
}

func TestDevice_GetByID_Device_Notfound(t *testing.T) {
	deviceServiceMock := interfaces.NewMockIDeviceService(t)
	deviceServiceMock.EXPECT().GetByID(mock.Anything).Return(nil, nil)

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)

	c.Params = []gin.Param{{Key: "id", Value: uuid.NewString()}}

	deviceHandler := Handler{
		deviceService: deviceServiceMock,
		logger:        logger,
	}

	deviceHandler.Get(c)
	assert.True(t, c.Writer.Status() == http.StatusNotFound)
}

func TestDevice_GetByID_Return_200_Ok(t *testing.T) {
	deviceServiceMock := interfaces.NewMockIDeviceService(t)
	deviceServiceMock.EXPECT().GetByID(mock.Anything).Return(
		&model.Device{
			DeviceID:  uuid.New(),
			Name:      uuid.NewString(),
			Brand:     uuid.NewString(),
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil)

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)

	c.Params = []gin.Param{{Key: "id", Value: uuid.NewString()}}

	deviceHandler := Handler{
		deviceService: deviceServiceMock,
		logger:        logger,
	}

	deviceHandler.Get(c)
	assert.True(t, c.Writer.Status() == http.StatusOK)
}

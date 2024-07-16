package device

import (
	"device-manager-service/cmd/api/device/contracts"
	service "device-manager-service/domain/device"
	deviceModel "device-manager-service/domain/device/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type DeviceHandler struct {
	deviceService *service.Device
}

const DEFAULT_LIMIT = 5
const DEFAULT_PAGE = 1

func NewDeviceHandler(deviceService *service.Device) *DeviceHandler {
	return &DeviceHandler{
		deviceService: deviceService,
	}
}

func (h DeviceHandler) Create(c *gin.Context) {
	var newDevice contracts.CreateDeviceRequest
	newDeviceCreated := time.Now()

	if err := c.BindJSON(&newDevice); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error Creating the Device %s", err.Error())})
		return
	}

	device := deviceModel.Device{
		DeviceID:  uuid.New(),
		Name:      newDevice.Name,
		Brand:     newDevice.Brand,
		IsActive:  true,
		CreatedAt: newDeviceCreated,
		UpdatedAt: newDeviceCreated,
	}

	err := h.deviceService.Create(&device)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Creating the Device %s", err.Error())})
		return
	}

	c.IndentedJSON(http.StatusCreated, device)
}

func (h DeviceHandler) Update(c *gin.Context) {
	var updatedDevice contracts.UpdateDeviceRequest
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid deviceID Parameter %s", err)})
		return
	}

	if err := c.BindJSON(&updatedDevice); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error Updating the Device %s", err.Error())})
		return
	}
	device, err := h.deviceService.GetByID(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Getting the Device %s", err.Error())})
		return
	}
	if device == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Device Not Found %s", deviceID)})
		return
	}

	device.Name = updatedDevice.Name
	device.Brand = updatedDevice.Brand
	device.IsActive = updatedDevice.IsActive

	err = h.deviceService.Update(device)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Updating the Device %s", err.Error())})
		return
	}

	c.IndentedJSON(http.StatusOK, device)
}

func (h DeviceHandler) Patch(c *gin.Context) {
	var patchDeviceBrand contracts.PatchDeviceRequest
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid deviceID Parameter %s", err)})
		return
	}

	if err := c.BindJSON(&patchDeviceBrand); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error Updating Brand the Device %s", err.Error())})
		return
	}

	device, err := h.deviceService.GetByID(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Getting the Device %s", err.Error())})
		return
	}
	if device == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Device Not Found %s", deviceID)})
		return
	}

	if !patchDeviceBrand.IsProcessable() {
		c.IndentedJSON(http.StatusNotModified, nil)
	}

	if patchDeviceBrand.Name != nil {
		device.Name = *patchDeviceBrand.Name
	}
	if patchDeviceBrand.Brand != nil {
		device.Brand = *patchDeviceBrand.Brand
	}
	if patchDeviceBrand.IsActive != nil {
		device.IsActive = *patchDeviceBrand.IsActive
	}

	err = h.deviceService.Update(device)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Updating Brand the Device %s", err.Error())})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (h DeviceHandler) Delete(c *gin.Context) {
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid deviceID Parameter %s", err)})
		return
	}

	err = h.deviceService.Delete(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Deleting the Device %s", err.Error())})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h DeviceHandler) Get(c *gin.Context) {
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid deviceID Parameter %s", err)})
		return
	}

	device, err := h.deviceService.GetByID(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Getting the Device %s", err.Error())})
		return
	}
	if device == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Device Not Found %s", deviceID)})
		return
	}

	c.IndentedJSON(http.StatusOK, device)
}

func (h DeviceHandler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(DEFAULT_LIMIT)))
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(DEFAULT_PAGE)))
	brand := c.DefaultQuery("brand", "")

	offset := limit * (page - 1)

	devices, err := h.deviceService.GetByAll(limit, offset, brand)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error Getting All Devices %s", err.Error())})
		return
	}

	c.IndentedJSON(http.StatusOK, devices)
}

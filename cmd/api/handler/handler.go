package handler

import (
	"database/sql"
	"device-manager-service/cmd/api/contracts"
	"device-manager-service/internal/device/interfaces"
	deviceModel "device-manager-service/internal/device/model"
	"device-manager-service/internal/device/repository"
	"device-manager-service/internal/device/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	deviceService interfaces.IDeviceService
	logger        *log.Logger
}

const DEFAULT_LIMIT = 5
const DEFAULT_PAGE = 1

func NewDeviceHandler(logger *log.Logger, db *sql.DB) *Handler {
	deviceRepository := repository.NewDeviceRepository(db)

	return &Handler{
		deviceService: service.NewDeviceService(logger, deviceRepository),
		logger:        logger,
	}
}

func (h Handler) Create(c *gin.Context) {
	var newDevice contracts.CreateDeviceRequest
	newDeviceCreated := time.Now()

	if err := c.BindJSON(&newDevice); err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("error Creating the IDeviceService", err))
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
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Creating the IDeviceService", err))
		return
	}

	c.IndentedJSON(http.StatusCreated, device)
}

func (h Handler) Update(c *gin.Context) {
	var updatedDevice contracts.UpdateDeviceRequest
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("invalid deviceID Parameter", err))
		return
	}

	if err := c.BindJSON(&updatedDevice); err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("error Updating the IDeviceService", err))
		return
	}
	device, err := h.deviceService.GetByID(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Getting the IDeviceService", err))
		return
	}
	if device == nil {
		c.IndentedJSON(http.StatusNotFound, h.buildReturnErrorMessage(fmt.Sprintf("IDeviceService Not Found %s", deviceID), nil))
		return
	}

	device.Name = updatedDevice.Name
	device.Brand = updatedDevice.Brand
	device.IsActive = updatedDevice.IsActive

	err = h.deviceService.Update(device)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Updating the IDeviceService", err))
		return
	}

	c.IndentedJSON(http.StatusOK, device)
}

func (h Handler) Patch(c *gin.Context) {
	var patchDeviceBrand contracts.PatchDeviceRequest
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("invalid deviceID Parameter", err))
		return
	}

	if err := c.BindJSON(&patchDeviceBrand); err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("error Updating Brand the IDeviceService", err))
		return
	}

	device, err := h.deviceService.GetByID(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Getting the IDeviceService", err))
		return
	}
	if device == nil {
		c.IndentedJSON(http.StatusNotFound, h.buildReturnErrorMessage(fmt.Sprintf("Device Not Found %s", deviceID), fmt.Errorf("device Not Found")))
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
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Updating Brand the IDeviceService", err))
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (h Handler) Delete(c *gin.Context) {
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("invalid deviceID Parameter", err))
		return
	}

	err = h.deviceService.Delete(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Deleting the IDeviceService", err))
		return
	}

	c.Status(http.StatusNoContent)
}

func (h Handler) Get(c *gin.Context) {
	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, h.buildReturnErrorMessage("invalid deviceID Parameter", err))
		return
	}

	device, err := h.deviceService.GetByID(deviceID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Getting the IDeviceService", err))
		return
	}
	if device == nil {
		c.IndentedJSON(http.StatusNotFound, h.buildReturnErrorMessage(fmt.Sprintf("Device Not Found %s", deviceID), fmt.Errorf("device Not Found")))
		return
	}

	c.IndentedJSON(http.StatusOK, device)
}

func (h Handler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(DEFAULT_LIMIT)))
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(DEFAULT_PAGE)))
	brand := c.DefaultQuery("brand", "")

	offset := limit * (page - 1)

	devices, err := h.deviceService.GetAll(limit, offset, brand)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, h.buildReturnErrorMessage("error Getting All Devices %s", err))
		return
	}

	c.IndentedJSON(http.StatusOK, devices)
}

func (h Handler) buildReturnErrorMessage(message string, err error) contracts.ErrorResponse {
	errorMessage := fmt.Sprintf("%s %s", message, err.Error())
	h.logger.Println(errorMessage)
	return contracts.ErrorResponse{
		Message: errorMessage,
	}
}

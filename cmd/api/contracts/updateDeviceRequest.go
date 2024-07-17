package contracts

type UpdateDeviceRequest struct {
	Name     string `json:"name" binding:"required"`
	Brand    string `json:"brand" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}

package contracts

type CreateDeviceRequest struct {
	Name  string `json:"name" binding:"required"`
	Brand string `json:"brand" binding:"required"`
}

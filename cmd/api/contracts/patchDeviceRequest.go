package contracts

type PatchDeviceRequest struct {
	Name     *string `json:"name,omitempty"`
	Brand    *string `json:"brand,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

func (p PatchDeviceRequest) IsProcessable() bool {
	return p.Name != nil || p.Brand != nil || p.IsActive != nil
}

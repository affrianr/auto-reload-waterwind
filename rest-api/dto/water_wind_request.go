package dto

type CreateWaterWindRequest struct {
	Water int `json:"water" binding:"required,min=1,max=100"`
	Wind  int `json:"wind" binding:"required,min=1,max=100"`
}

type UpdateWaterWindRequest struct {
	Water int `json:"water" binding:"required,min=1,max=100"`
	Wind  int `json:"wind" binding:"required,min=1,max=100"`
}

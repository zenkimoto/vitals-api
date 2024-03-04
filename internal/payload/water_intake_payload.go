package payload

import (
	"time"

	"github.com/zenkimoto/vitals-server-api/internal/models"
)

type WaterIntakeRequest struct {
	Cups float32   `json:"cups" binding:"required"`
	Time time.Time `json:"time"`
}

type WaterIntakeResponse struct {
	Id   uint      `json:"id" binding:"required"`
	Cups float32   `json:"cups" binding:"required"`
	Time time.Time `json:"time"`
}

func MapWaterIntakeResponse(wi models.WaterIntake) WaterIntakeResponse {
	return WaterIntakeResponse{
		Id:   wi.ID,
		Cups: wi.Cups,
		Time: wi.Time,
	}
}

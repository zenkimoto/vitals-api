package payload

import (
	"time"

	"github.com/zenkimoto/vitals-server-api/internal/models"
)

type WeightRequest struct {
	Weight float32   `json:"weight" binding:"required"`
	Time   time.Time `json:"time"`
}

type WeightResponse struct {
	Id     uint      `json:"id"`
	Weight float32   `json:"weight"`
	Time   time.Time `json:"time"`
}

func MapWeightResponse(w models.Weight) WeightResponse {
	return WeightResponse{
		Id:     w.ID,
		Weight: w.Weight,
		Time:   w.Time,
	}
}

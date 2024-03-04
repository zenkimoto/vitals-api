package payload

import (
	"time"

	"github.com/zenkimoto/vitals-server-api/internal/models"
)

type BloodPressureRequest struct {
	Sys  uint16    `json:"systolic" binding:"required"`
	Dia  uint16    `json:"diastolic" binding:"required"`
	Time time.Time `json:"time"`
}

type BloodPressureResponse struct {
	Id   uint      `json:"id"`
	Sys  uint16    `json:"systolic"`
	Dia  uint16    `json:"diastolic"`
	Time time.Time `json:"time"`
}

func MapBloodPressureResponse(bp models.BloodPressure) BloodPressureResponse {
	return BloodPressureResponse{
		Id:   bp.ID,
		Sys:  bp.Sys,
		Dia:  bp.Dia,
		Time: bp.Time,
	}
}

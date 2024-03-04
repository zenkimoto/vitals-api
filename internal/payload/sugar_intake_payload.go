package payload

import (
	"time"

	"github.com/zenkimoto/vitals-server-api/internal/models"
)

type SugarIntakeRequest struct {
	Grams uint      `json:"grams" binding:"required"`
	Time  time.Time `json:"time"`
}

type SugarIntakeResponse struct {
	Id    uint      `json:"id" binding:"required"`
	Grams uint      `json:"grams" binding:"required"`
	Time  time.Time `json:"time"`
}

func MapSugarIntakeResponse(si models.SugarIntake) SugarIntakeResponse {
	return SugarIntakeResponse{
		Id:    si.ID,
		Grams: si.Grams,
		Time:  si.Time,
	}
}

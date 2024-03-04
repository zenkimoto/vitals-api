package payload

import (
	"time"

	"github.com/zenkimoto/vitals-server-api/internal/models"
)

type UserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
	UserName  string    `json:"username"`
	UserSince time.Time `json:"userSince"`
}

func MapUserResponse(u models.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		UserName:  u.UserName,
		UserSince: u.CreatedAt,
	}
}

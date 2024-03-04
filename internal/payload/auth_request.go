package payload

// Credentials Request payload
type AuthRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

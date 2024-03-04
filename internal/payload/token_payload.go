package payload

// Token Request payload
type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// Token Response payload
type TokenResponse struct {
	Token string `json:"token" binding:"required"`
}

type ValidateTokenResponse struct {
	UserId   uint   `json:"id" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

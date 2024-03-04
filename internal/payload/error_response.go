package payload

// Error Response payload
type ErrorResponse struct {
	Error string `json:"error" binding:"required"`
}

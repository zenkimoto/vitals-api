package payload

// HealthCheckResponse serves as the response body for the HealthCheck endpoint
type HealthCheckResponse struct {
	Status string `json:"status"`
}

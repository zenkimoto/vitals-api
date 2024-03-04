package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
)

// HealthCheck()
// If the server is up and running, the endpoint will return a JSON object with a status of "ok".
// GET /health-check
//
// Swagger Doc
// @Summary Health Check
// @Schemes
// @Description If the server is up and running, the endpoint will return a JSON object with a status of "ok".
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} payload.HealthCheckResponse
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, payload.HealthCheckResponse{Status: "ok"})
}

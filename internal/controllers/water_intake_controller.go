package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/models"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
)

// GET /users/:id/water
// Get all water intake records for a user.
//
// Swagger Doc
// @Summary Get all water intake records for a user.
// @Schemes
// @Description Get all water intake records for a user.
// @Tags Water Intake
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} payload.WaterIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/water [get]
// @Security Bearer
func GetWaterIntakeByUserId(c *gin.Context) {
	userId := c.Param("id")
	var wi []models.WaterIntake

	var u models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "User not found"})
		return
	}

	models.DB.Where("user_id = ?", userId).Order("time DESC").Find(&wi)

	c.JSON(http.StatusOK, Map(wi, payload.MapWaterIntakeResponse))
}

// POST /users/:id/water
// Adds a new water intake record for user.
//
// Swagger Doc
// @Summary Adds a new water intake record for user.
// @Schemes
// @Description Adds a new water intake record for user. Time is optional. If not provided, current time is used.
// @Tags Water Intake
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body payload.WaterIntakeRequest true "User Water Intake"
// @Success 200 {object} payload.WaterIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/water [post]
// @Security Bearer
func PostWaterIntakeByUserId(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	// Validate r
	var r payload.WaterIntakeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	createTime := time.Now()
	if r.Time.IsZero() == false {
		createTime = r.Time
	}

	// Create Water Intake Record
	wi := models.WaterIntake{Cups: r.Cups, UserID: uint(id), Time: createTime}
	models.DB.Create(&wi)

	c.JSON(http.StatusOK, payload.MapWaterIntakeResponse(wi))
}

// PUT /users/:userId/water/:id
// Updates a water intake record for user.
//
// Swagger Doc
// @Summary Updates a water intake record for user.
// @Schemes
// @Description Updates a water intake record for user. Time is optional.
// @Tags Water Intake
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Water Intake ID"
// @Param user body payload.WaterIntakeRequest true "User Water Intake"
// @Success 200 {object} payload.WaterIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/water/{id} [put]
// @Security Bearer
func PutWaterIntakeByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid water intake id"})
		return
	}

	// Validate r
	var r payload.WaterIntakeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	// Find Water Intake Record
	var wi models.WaterIntake
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&wi).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Water Intake not found"})
		return
	}

	// Update Water Intake Record
	wi.Cups = r.Cups
	models.DB.Save(&wi)

	c.JSON(http.StatusOK, payload.MapWaterIntakeResponse(wi))
}

// DELETE /users/:userId/water/:id
// Deletes a water intake record for user.
//
// Swagger Doc
// @Summary Deletes a water intake record for user.
// @Schemes
// @Description Deletes a water intake record for user.
// @Tags Water Intake
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Water Intake ID"
// @Success 200 {object} payload.WaterIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/water/{id} [delete]
// @Security Bearer
func DeleteWaterIntakeByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid water id"})
		return
	}

	// Delete Water Intake Record
	var wi models.WaterIntake
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&wi).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Water Intake not found"})
		return
	}

	models.DB.Delete(&wi)

	c.JSON(http.StatusOK, payload.MapWaterIntakeResponse(wi))
}

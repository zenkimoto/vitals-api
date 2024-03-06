package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/models"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
)

// GET /users/:id/blood-pressure
// Get all blood pressure records for a user.
//
// Swagger Doc
// @Summary Get all blood pressure records for a user.
// @Schemes
// @Description Get all blood pressure records for a user.
// @Tags Blood Pressure
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} payload.BloodPressureResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/blood-pressure [get]
// @Security Bearer
func GetBloodPressureByUserId(c *gin.Context) {
	userId := c.Param("id")
	var bp []models.BloodPressure

	var u models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "User not found"})
		return
	}

	models.DB.Where("user_id = ?", userId).Order("time DESC").Find(&bp)

	c.JSON(http.StatusOK, Map(bp, payload.MapBloodPressureResponse))
}

// POST /users/:id/blood-pressure
// Adds a new blood pressure record for user.
//
// Swagger Doc
// @Summary Adds a new blood pressure record for user.
// @Schemes
// @Description Adds a new blood pressure record for user. Time is optional. If not provided, current time is used.
// @Tags Blood Pressure
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body payload.BloodPressureRequest true "User Blood Pressure"
// @Success 200 {object} payload.BloodPressureResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/blood-pressure [post]
// @Security Bearer
func PostBloodPressureByUserId(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	// Validate r
	var r payload.BloodPressureRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	createTime := time.Now()
	if r.Time.IsZero() == false {
		createTime = r.Time
	}

	// Create Blood Pressure
	bp := models.BloodPressure{Sys: r.Sys, Dia: r.Dia, UserID: uint(id), Time: createTime}
	models.DB.Create(&bp)

	c.JSON(http.StatusOK, payload.MapBloodPressureResponse(bp))
}

// PUT /users/:userId/blood-pressure/:id
// Updates a blood pressure record for user.
//
// Swagger Doc
// @Summary Updates a blood pressure record for user.
// @Schemes
// @Description Updates a blood pressure record for user. Time is optional.
// @Tags Blood Pressure
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Blood Pressure ID"
// @Param user body payload.BloodPressureRequest true "User Blood Pressure"
// @Success 200 {object} payload.BloodPressureResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/blood-pressure/{id} [put]
// @Security Bearer
func PutBloodPressureByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid blood pressure id"})
		return
	}

	// Validate r
	var r payload.BloodPressureRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	// Find Blood Pressure
	var bp models.BloodPressure
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&bp).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Blood Pressure not found"})
		return
	}

	// Update Blood Pressure
	bp.Sys = r.Sys
	bp.Dia = r.Dia
	models.DB.Save(&bp)

	c.JSON(http.StatusOK, payload.MapBloodPressureResponse(bp))
}

// DELETE /users/:userId/blood-pressure/:id
// Deletes a blood pressure record for user.
//
// Swagger Doc
// @Summary Deletes a blood pressure record for user.
// @Schemes
// @Description Deletes a blood pressure record for user.
// @Tags Blood Pressure
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Blood Pressure ID"
// @Success 200 {object} payload.BloodPressureResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/blood-pressure/{id} [delete]
// @Security Bearer
func DeleteBloodPressureByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid blood pressure id"})
		return
	}

	// Find Blood Pressure
	var bp models.BloodPressure
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&bp).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Blood Pressure not found"})
		return
	}

	// Delete Blood Pressure
	models.DB.Delete(&bp)

	c.JSON(http.StatusOK, payload.MapBloodPressureResponse(bp))
}

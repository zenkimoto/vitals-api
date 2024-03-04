package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/models"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
)

// GET /users/:id/sugar
// Get all sugar intake records for a user.
//
// Swagger Doc
// @Summary Get all sugar intake records for a user.
// @Schemes
// @Description Get all sugar intake records for a user.
// @Tags Sugar Intake
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} payload.SugarIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/sugar [get]
// @Security Bearer
func GetSugarIntakeByUserId(c *gin.Context) {
	var u models.User
	if err := models.DB.Where("id = ?", c.Param("id")).Preload("SugarIntakeList").First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, Map(u.SugarIntakeList, payload.MapSugarIntakeResponse))
}

// POST /users/:id/sugar
// Adds a new sugar intake record for user.
//
// Swagger Doc
// @Summary Adds a new sugar intake record for user.
// @Schemes
// @Description Adds a new sugar intake record for user. Time is optional. If not provided, current time is used.
// @Tags Sugar Intake
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body payload.SugarIntakeRequest true "User Sugar Intake"
// @Success 200 {object} payload.SugarIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/sugar [post]
// @Security Bearer
func PostSugarIntakeByUserId(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	// Validate r
	var r payload.SugarIntakeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	createTime := time.Now()
	if r.Time.IsZero() == false {
		createTime = r.Time
	}

	// Create Sugar Intake Record
	si := models.SugarIntake{Grams: r.Grams, UserID: uint(id), Time: createTime}
	models.DB.Create(&si)

	c.JSON(http.StatusOK, payload.MapSugarIntakeResponse(si))
}

// PUT /users/:userId/sugar/:id
// Updates a sugar intake record for user
//
// Swagger Doc
// @Summary Updates a sugar intake record for user
// @Schemes
// @Description Updates a sugar intake record for user
// @Tags Sugar Intake
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Sugar Intake ID"
// @Param user body payload.SugarIntakeRequest true "User Sugar Intake"
// @Success 200 {object} payload.SugarIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/sugar/{id} [put]
// @Security Bearer
func PutSugarIntakeByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid sugar intake id"})
		return
	}

	// Validate Request
	var r payload.SugarIntakeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	// Find Sugar Intake Record
	var si models.SugarIntake
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&si).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Sugar Intake not found"})
		return
	}

	// Update Sugar Intake Record
	si.Grams = r.Grams
	models.DB.Save(&si)

	c.JSON(http.StatusOK, payload.MapSugarIntakeResponse(si))
}

// DELETE /users/:userId/sugar/:id
// Deletes a sugar intake record for user.
//
// Swagger Doc
// @Summary Deletes a sugar intake record for user.
// @Schemes
// @Description Deletes a sugar intake record for user.
// @Tags Sugar Intake
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Sugar Intake ID"
// @Success 200 {object} payload.SugarIntakeResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/sugar/{id} [delete]
// @Security Bearer
func DeleteSugarIntakeByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid sugar id"})
		return
	}

	// Find Sugar Intake Record
	var si models.SugarIntake
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&si).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Sugar Intake not found"})
		return
	}

	// Delete Sugar Intake Record
	models.DB.Delete(&si)

	c.JSON(http.StatusOK, payload.MapSugarIntakeResponse(si))
}

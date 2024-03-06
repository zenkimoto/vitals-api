package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/models"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
)

// GET /users/:id/weight
// Get all weight records for a user.
//
// Swagger Doc
// @Summary Get all weight records for a user.
// @Schemes
// @Description Get all weight records for a user.
// @Tags Weight
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} payload.WeightResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/weight [get]
// @Security Bearer
func GetWeightByUserId(c *gin.Context) {
	userId := c.Param("id")
	var w []models.Weight

	var u models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "User not found"})
		return
	}

	models.DB.Where("user_id = ?", userId).Order("time DESC").Find(&w)

	c.JSON(http.StatusOK, Map(w, payload.MapWeightResponse))
}

// POST /users/:id/weight
// Adds a new weight record for user
//
// Swagger Doc
// @Summary Adds a new weight record for user.
// @Schemes
// @Description Adds a new weight record for user. Time is optional. If not provided, current time is used.
// @Tags Weight
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body payload.WeightRequest true "User Weight"
// @Success 200 {object} payload.WeightResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id}/weight [post]
// @Security Bearer
func PostWeightByUserId(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	// Validate Request
	var request payload.WeightRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	createTime := time.Now()
	if request.Time.IsZero() == false {
		createTime = request.Time
	}

	// Create Weight
	w := models.Weight{Weight: request.Weight, UserID: uint(id), Time: createTime}
	models.DB.Create(&w)

	c.JSON(http.StatusOK, payload.MapWeightResponse(w))
}

// PUT /users/:userId/weight/:id
// Updates a weight record for user.
//
// Swagger Doc
// @Summary Updates a weight record for user.
// @Schemes
// @Description Updates a weight record for user. Time is optional.
// @Tags Weight
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Weight ID"
// @Param user body payload.WeightRequest true "User Weight"
// @Success 200 {object} payload.WeightResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/weight/{id} [put]
// @Security Bearer
func PutWeightByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid weight id"})
		return
	}

	// Validate Request
	var req payload.WeightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	// Find Weight
	var w models.Weight
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&w).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Weight not found"})
		return
	}

	// Update Weight
	w.Weight = req.Weight
	models.DB.Save(&w)

	c.JSON(http.StatusOK, payload.MapWeightResponse(w))
}

// DELETE /users/:userId/weight/:id
// Deletes a weight record for user.
//
// Swagger Doc
// @Summary Deletes a weight record for user.
// @Schemes
// @Description Deletes a weight record for user.
// @Tags Weight
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param id path int true "Weight ID"
// @Success 200 {object} payload.WeightResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{userId}/weight/{id} [delete]
// @Security Bearer
func DeleteWeightByUserId(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid user id"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid weight id"})
		return
	}

	// Find Weight
	var w models.Weight
	if err := models.DB.Where("id = ? AND user_id = ?", id, userId).First(&w).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "Weight not found"})
		return
	}

	// Delete Weight
	models.DB.Delete(&w)

	c.JSON(http.StatusOK, payload.MapWeightResponse(w))
}

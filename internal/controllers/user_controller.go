package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/models"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
)

// GET /users
// Get all users
//
// Swagger Doc
// @Summary Get all users
// @Schemes
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} payload.UserResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users [get]
// @Security Bearer
func GetUsers(c *gin.Context) {
	var userList []models.User
	models.DB.Find(&userList)

	c.JSON(http.StatusOK, Map(userList, payload.MapUserResponse))
}

// GET /users/:id
// Get user by user id
//
// Swagger Doc
// @Summary Get user by user id
// @Schemes
// @Description Get user by user id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} payload.UserResponse
// @Failure 404 {object} payload.ErrorResponse
// @Router /users/{id} [get]
// @Security Bearer
func GetUserById(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, payload.ErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, payload.MapUserResponse(user))
}

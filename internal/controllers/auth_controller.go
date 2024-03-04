package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/env"
	"github.com/zenkimoto/vitals-server-api/internal/models"
	"github.com/zenkimoto/vitals-server-api/internal/payload"
	"github.com/zenkimoto/vitals-server-api/internal/util"
)

// Login POST /auth/login
// Login request handler reads the username and password from the request body
// and checks if the user exists in the database. If the user exists, the
// password is verified using bcrypt. If the password is verified, a JWT is
// issued and returned to the client.
//
// Swagger Doc
// @Summary Login to the Vital Server API
// @Schemes
// @Description Authenticates a user with username and password credentials. If the credentials are valid, the endpoint will issue a JSON Web Token (JWT) to the client for the duration of the session.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body payload.AuthRequest true "User Credentials Request"
// @Success 200 {object} payload.TokenResponse
// @Failure 400 {object} payload.ErrorResponse
// @Router /auth [post]
func Login(c *gin.Context) {
	var request payload.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	// Check if user exists
	var user models.User
	if err := models.DB.Where("user_name = ?", request.UserName).First(&user).Error; err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid username and/or password"})
		return
	}

	// Validate bcrypt password hash
	if !util.VerifyPassword(request.Password, user.PasswordHash) {
		log.Print("bcrypt hash does not match.")
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid username and/or password"})
		return
	}

	jwt, err := issueJsonWebToken(user)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{Error: "Internal Server Error"})
	}

	c.JSON(http.StatusOK, payload.TokenResponse{Token: jwt})
}

// Issues a JSON Web Token.  The token is signed with the JWT key (set
// by an env var) and contains the user's username and id.
// The token expires after the duration specified in the env vars.
func issueJsonWebToken(user models.User) (string, error) {
	key := env.GetJWTKey()
	duration := env.GetTokenExpirationDuration()

	return util.Issue(key, user.UserName, user.ID, duration)
}

// ValidateToken POST /token/validate
// Validates a JSON Web Token.  If the token is valid, it returns the
// user and id associated with the token.
//
// @Summary Validates a JSON Web Token
// @Schemes
// @Description Verifies the validity of a JSON Web Token (JWT) and returns the user and id associated with the token. Can be used by the client to determine if a token needs to be refreshed.
// @Tags Token
// @Accept json
// @Produce json
// @Param user body payload.TokenRequest true "Token"
// @Success 200 {object} payload.ValidateTokenResponse
// @Failure 400 {object} payload.ErrorResponse
// @Router /token/validate [post]
func ValidateToken(c *gin.Context) {
	var request payload.TokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	key := env.GetJWTKey()
	username, id, err := util.Parse(key, request.Token)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, payload.ValidateTokenResponse{UserId: id, UserName: username})
}

// RefreshToken POST /token/refresh
// Refreshes a JSON Web Token.  If the token is valid, it returns a new
// token with the same user and id with a new expiration.
//
// @Summary Refreshes a validated JSON Web Token
// @Schemes
// @Description Verifies the validity of a JSON Web Token (JWT) and if valid, the endpoint will issue a new JSON Web Token.
// @Tags Token
// @Accept json
// @Produce json
// @Param user body payload.TokenRequest true "Token"
// @Success 200 {object} payload.TokenResponse
// @Failure 400 {object} payload.ErrorResponse
// @Router /token/refresh [post]
func RefreshToken(c *gin.Context) {
	var request payload.TokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: err.Error()})
		return
	}

	key := env.GetJWTKey()
	username, id, err := util.Parse(key, request.Token)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid token"})
		return
	}

	var user models.User
	err = models.DB.Where("id = ? AND user_name = ?", id, username).First(&user).Error

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{Error: "Invalid token"})
		return
	}

	jwt, err := issueJsonWebToken(user)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{Error: "Internal Server Error"})
	}

	c.JSON(http.StatusOK, payload.TokenResponse{Token: jwt})
}

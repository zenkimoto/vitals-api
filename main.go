package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	docs "github.com/zenkimoto/vitals-server-api/docs"
	"github.com/zenkimoto/vitals-server-api/internal/controllers"
	"github.com/zenkimoto/vitals-server-api/internal/env"
	"github.com/zenkimoto/vitals-server-api/internal/middleware"
	"github.com/zenkimoto/vitals-server-api/internal/models"
)

func main() {
	env.LoadEnvironmentVariables()

	host, user, password, dbname := env.GetDatabaseConnectionInfo()

	models.InitializeDatabase(host, user, password, dbname)

	startServer()
}

// @title           Vitals Server API
// @version         1.0
// @description     <h3>Vitals API is a simple API for tracking health vitals and lifestyle.</h3>
// @description     <p>The Vitals API tracks weight, blood pressure, water and sugar intake.</p>
// @description		<h4>To Use the Vitals API:</h4>
// @description     <ol>
// @description     <p><li>Log into the /auth endpoint.</li></p>
// @description     <p><li>Once successful, you will get a receive a token in the authentication response.<p>The token must be added in the Authorization header in any of the secured endpoints.</p><p>Authorization: Bearer {token}</p></li></p>
// @description     <p><li>Call any of the endpoints: /weight, /blood-pressure, /water, /sugar</li></p>
// @description     </ol>
//
// @contact.name	Alex Yip
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func startServer() {
	router := gin.Default()
	port := env.GetPort()

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Public Routes
	router.GET("/health-check", controllers.HealthCheck)

	router.POST("/auth", controllers.Login)
	router.POST("/token/validate", controllers.ValidateToken)
	router.POST("/token/refresh", controllers.RefreshToken)

	// Swagger Set Up
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// JWT Auth Middleware
	router.Use(middleware.JwtAuth())

	// Protected Routes
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserById)

	router.GET("/users/:id/blood-pressure", controllers.GetBloodPressureByUserId)
	router.POST("/users/:id/blood-pressure", controllers.PostBloodPressureByUserId)
	router.PUT("/users/:userId/blood-pressure/:id", controllers.PutBloodPressureByUserId)
	router.DELETE("/users/:userId/blood-pressure/:id", controllers.DeleteBloodPressureByUserId)

	router.GET("/users/:id/weight", controllers.GetWeightByUserId)
	router.POST("/users/:id/weight", controllers.PostWeightByUserId)
	router.PUT("/users/:userId/weight/:id", controllers.PutWeightByUserId)
	router.DELETE("/users/:userId/weight/:id", controllers.DeleteWeightByUserId)

	router.GET("/users/:id/sugar", controllers.GetSugarIntakeByUserId)
	router.POST("/users/:id/sugar", controllers.PostSugarIntakeByUserId)
	router.PUT("/users/:userId/sugar/:id", controllers.PutSugarIntakeByUserId)
	router.DELETE("/users/:userId/sugar/:id", controllers.DeleteSugarIntakeByUserId)

	router.GET("/users/:id/water", controllers.GetWaterIntakeByUserId)
	router.POST("/users/:id/water", controllers.PostWaterIntakeByUserId)
	router.PUT("/users/:userId/water/:id", controllers.PutWaterIntakeByUserId)
	router.DELETE("/users/:userId/water/:id", controllers.DeleteWaterIntakeByUserId)

	router.Run(port)
}

package routes

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"secure-login-system/handlers"
	"secure-login-system/middleware"
)

func SetupRoutes(db *sql.DB, jwtSecret string) *gin.Engine {
	authHandler := &handlers.AuthHandler{
		DB:        db,
		JWTSecret: jwtSecret,
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/api/login", authHandler.Login)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(jwtSecret))
	protected.GET("/dashboard", handlers.Dashboard)

	return router
}

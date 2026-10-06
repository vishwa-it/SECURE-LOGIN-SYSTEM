package main

import (
	"log"

	"secure-login-system/config"
	"secure-login-system/database"
	"secure-login-system/routes"
)

func main() {
	cfg := config.Load()

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = "8080"
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	log.Println("Database connected successfully!")

	router := routes.SetupRoutes(db, cfg.JWTSecret)

	log.Printf("Server running on http://localhost:%s", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Server failed:", err)
	}
}

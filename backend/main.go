// backend/main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/config"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/controllers"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/middlewares"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/routes"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/utils"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize JWT utility
    jwtUtil := utils.NewJWTUtil(cfg.JWTSecret)

    // Initialize database
    db, err := config.InitDB(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer config.CloseDB(db)

    // Initialize service
    userService := controllers.NewUserService(db, jwtUtil)

    // Initialize controller
    userController := controllers.NewUserController(userService)

    // Initialize Gin router
    router := gin.Default()

    // Apply middleware
    router.Use(middlewares.Logger())

    // Setup routes
    routes.SetupRoutes(router, userController, jwtUtil)

    // Start server
    if err := router.Run(":" + cfg.Port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

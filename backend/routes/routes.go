// backend/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/controllers"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/middlewares"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/utils"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController, jwtUtil *utils.JWTUtil) {
    // Public routes
    public := router.Group("/api")
    {
        public.POST("/register", userController.Register)
        public.POST("/login", userController.Login)
    }

    // Protected routes
    protected := router.Group("/api")
    protected.Use(middlewares.AuthMiddleware(jwtUtil))
    {
        protected.GET("/users/:id", userController.GetUser)
        // Tambahkan route lain yang memerlukan autentikasi
    }
}

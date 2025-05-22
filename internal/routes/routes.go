package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/AIhmed/go-api-test/internal/config"
	"github.com/AIhmed/go-api-test/internal/controllers"
	"github.com/AIhmed/go-api-test/internal/middleware"
	"github.com/AIhmed/go-api-test/internal/repositories"
	"github.com/AIhmed/go-api-test/internal/services"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)

	// Middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())

	// API routes
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.POST("/", userController.CreateUser)
				// Add other user routes
			}
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}

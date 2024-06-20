package routes

import (
	"github.com/iffakhry/go-commerce-mvc/controller"
	"github.com/iffakhry/go-commerce-mvc/pkg/middlewares"
	"github.com/iffakhry/go-commerce-mvc/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} \n",
	}))

	userRepo := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepo)
	e.POST("/login", userController.Login)

	e.GET("/profile", userController.GetProfile, middlewares.JWTMiddleware())
	e.GET("/users", userController.GetAllUser, middlewares.JWTMiddleware())
	e.POST("/users", userController.CreateUser)
	e.PUT("/users", userController.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userController.Delete, middlewares.JWTMiddleware())

	// e.POST("/users/alamat", controllers.AddUserAlamatController)

	// product := e.Group("/products", middlewares.JWTMiddleware())
	// product.GET("", controllers.GetProductController)
	// product.POST("", controllers.AddProductController)
	// return e
}

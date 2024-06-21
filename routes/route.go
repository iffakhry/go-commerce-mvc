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
	api := e.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.POST("/login", userController.Login)

	apiV1.GET("/profile", userController.GetProfile, middlewares.JWTMiddleware())
	apiV1.GET("/users", userController.GetAllUser, middlewares.JWTMiddleware())
	apiV1.POST("/users", userController.CreateUser)
	apiV1.PUT("/users", userController.Update, middlewares.JWTMiddleware())
	apiV1.DELETE("/users", userController.Delete, middlewares.JWTMiddleware())

	productV1 := apiV1.Group("/products")

	productRepo := repository.NewProductRepository(db)
	productController := controller.NewProductController(productRepo)
	productV1.POST("", productController.Create, middlewares.JWTMiddleware())
	productV1.GET("", productController.GetAll)
	productV1.GET("/:id", productController.GetById)
	productV1.PUT("/:id", productController.Update, middlewares.JWTMiddleware())
	productV1.DELETE("/:id", productController.Delete, middlewares.JWTMiddleware())
	// return e
}

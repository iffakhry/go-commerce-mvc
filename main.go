package main

import (
	"fmt"

	"github.com/iffakhry/go-commerce-mvc/config"
	_ "github.com/iffakhry/go-commerce-mvc/docs"
	"github.com/iffakhry/go-commerce-mvc/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Go-Commerce API
// @version 1.0
// @description This is API documentation for Go-Commerce.
// @termsOfService http://swagger.io/terms/

// @contact.name Fakhry Firdaus
// @contact.url http://academy.alterra.id
// @contact.email fakhry@alterra.id

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8080
// @BasePath /api/v1
func main() {
	fmt.Println("hello world")
	cfg := config.InitConfig()
	db := config.InitDBPostgres(cfg)
	config.InitMigrate(db)

	// create new instance echo
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.InitRoute(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}

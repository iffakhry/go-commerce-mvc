package main

import (
	"fmt"

	"github.com/iffakhry/go-commerce-mvc/config"
	"github.com/iffakhry/go-commerce-mvc/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello world")
	cfg := config.InitConfig()
	db := config.InitDBPostgres(cfg)
	config.InitMigrate(db)

	// create new instance echo
	e := echo.New()
	routes.InitRoute(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}

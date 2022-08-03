package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common"
	"main/common/db"
)

func main() {
	if err := db.InitMongo(); err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	if err := common.InitHandler(e); err != nil {
		fmt.Println(err)
	}

	e.Logger.Fatal(e.Start(":3000"))
}

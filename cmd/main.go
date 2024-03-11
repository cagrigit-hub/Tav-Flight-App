package main

import (
	"github.com/cagrigit-hub/tav-app/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	homeHandler := &handler.HomeHandler{}
	app.Static("/static", "assets")
	app.GET("/", homeHandler.HandleHomeShow)
	app.POST("/upload-excel", homeHandler.HandleExcelPost)
	app.Start(":3000")
}

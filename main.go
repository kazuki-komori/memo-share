package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", hoge)
	e.Logger.Fatal(e.Start(":8080"))
}

func hoge(c echo.Context) error {
	return c.String(200, "fuga")
}

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

)

func main() {
	log.Println("start")
	e := echo.New()

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi")
}

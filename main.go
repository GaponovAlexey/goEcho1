package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

)

func main() {
	port := os.Getenv("MY_PORT")
	log.Println("start")
	e := echo.New()

	e.GET("/", hello)
	e.Logger.Print("test")
	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi")
}

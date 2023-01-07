package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

)

func main() {
	log.Println("start")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//
	port := os.Getenv("MY_PORT") // "3000"
	log.Println(port)

	//start
	e := echo.New()
	//func
	e.GET("/", hello)
	//end
	e.Logger.Fatal(e.Start(port))
}

// func
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi")
}

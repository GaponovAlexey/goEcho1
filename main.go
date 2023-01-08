package main

import (
	"log"
	h "net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

)

func main() {
	log.Println("start")
	//install .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//get port .env
	port := os.Getenv("MY_PORT") // "3000"
	log.Println(port)

	//validator
	v := validator.New()
	log.Println(v)

	products := []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptop"}}

	//start
	e := echo.New()
	//func
	e.GET("/", hello)

	e.GET("/p", func(c echo.Context) error {
		return c.JSON(h.StatusOK, products)
	})

	e.GET("/p/:id", func(c echo.Context) error {
		var product map[int]string

		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(h.StatusNotFound, "product not found")
		}
		return c.JSON(h.StatusOK, product)
	})

	e.POST("/p", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil {
			return err
		}
		// product := map[int]string{
			// len(products) + 1: reqBody.Name,
		// }

	})
	//end
	e.Logger.Fatal(e.Start(port))
}

// func
func hello(c echo.Context) error {
	return c.String(h.StatusOK, "hi")
}

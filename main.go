package main

import (
	"log"
	h "net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

)

var (
	products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptop"}}
	v        = validator.New()
	e        = echo.New()
)

type ProdVal struct {
	val *validator.Validate
}

func (p *ProdVal) Validate(i interface{}) error {
	return p.val.Struct(i)
}

func main() {
	log.Println("start")

	//func
	e.GET("/", hello)

	e.GET("/p", func(c echo.Context) error {
		return c.JSON(h.StatusOK, products)
	})
	//getId -- d
	e.GET("/prod/:id", getId)
	//POST
	e.POST("/p", addProd)

	//end
	e.Logger.Fatal(e.Start(":3000"))
}

// func
func hello(c echo.Context) error {
	return c.String(h.StatusOK, "hi")
}

func getId(c echo.Context) error {
	paragraphID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}
	var newProduct map[int]string
	for _, v := range products {
		for k := range v {
			if k == paragraphID {
				newProduct = v
			}
		}
	}
	if newProduct == nil {
		return c.JSON(h.StatusNotFound, "not found")
	}

	return c.JSON(h.StatusOK, newProduct)
}

func addProd(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProdVal{val: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(h.StatusOK, product)
}

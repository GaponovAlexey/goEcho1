package main

import (
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
	//func
	e.GET("/p", func(c echo.Context) error {
		return c.JSON(h.StatusOK, products)
	})
	//getId -- id
	e.GET("/p/:id", getId)
	//POST
	e.POST("/p", addProd)
	e.PUT("/p/:id", putProd)
	e.DELETE("/p/:id", delProd)

	//end
	e.Logger.Fatal(e.Start(":3000"))
}

// func
func getId(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	var newProduct map[int]string

	for _, v := range products {
		for k := range v {
			if pID == k {
				newProduct = v
			}
		}
	}
	if newProduct == nil {
		return c.JSON(h.StatusNotFound, "not found product id")
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

func putProd(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("id"))
	var prod map[int]string

	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pID == k {
				prod = p
			}
		}
	}

	if prod == nil {
		return c.JSON(h.StatusNotFound, "not data")
	}

	type req struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}

	var body req

	e.Validator = &ProdVal{val: v}

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return err
	}

	prod[pID] = body.Name
	return c.JSON(h.StatusOK, prod)
}

func delProd(c echo.Context) error {
	var prod map[int]string
	var index int
	pID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	for i, p := range products {
		for k := range p {
			if pID == k {
				prod = p
				index = i
			}
		}
	}
	if prod == nil {
		return c.JSON(h.StatusNotFound, "not Found")
	}
	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}

	products = splice(products, index)

	return c.JSON(h.StatusOK, prod)
}


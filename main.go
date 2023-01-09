package main

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

)

type Valid struct {
	Valid *validator.Validate
}
type req struct {
	// Name string `json:"name" validate:"required,min=4"`
	Name string `json:"name"`
}

var (
	mas    = []map[int]string{{1: "2"}, {2: "two"}, {3: "three"}}
	e      = echo.New()
	vNew   = validator.New()
	body   = req{}
	newMas = map[int]string{}
)

func (p *Valid) Validate(i interface{}) error {
	return p.Valid.Struct(i)
}

func main() {

	//start
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, mas)
	})
	e.GET("/:id", getId)
	e.POST("/", postMas)
	e.PUT("/:id", putMas)
	e.DELETE("/:id", delMas)
	//end
	e.Logger.Fatal(e.Start(":3000"))
}

func cancel(e error) {
	if e != nil {
		panic(e)
	}
}

// func
func putMas(c echo.Context) error {
	pId, err := strconv.Atoi(c.Param("id"))
	cancel(err)

	if err := c.Bind(&body); err != nil {
		return err
	}
	for _, v := range mas {
		for k := range v {
			if k == pId {
				newMas = v
			}
		}
	}
	newMas[pId] = body.Name
	return c.JSON(http.StatusOK, newMas)
}

func delMas(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("id"))
	cancel(err)
	var index int
	for i, v := range mas {
		for k := range v {
			if k == pID {
				newMas = v
				index = i
			}
		}
	}
	if newMas == nil {
		return c.JSON(http.StatusNotFound, "not found")
	}
	splice := func(s []map[int]string, i int) []map[int]string {
		return append(s[:i], s[i+1:]...)
	}
	mas = splice(mas, index)
	return c.JSON(http.StatusOK, newMas)
}

func getId(c echo.Context) error {
	var prod map[int]string
	gId, err := strconv.Atoi(c.Param("id"))
	cancel(err)
	for _, v := range mas {
		for k := range v {
			if k == gId {
				prod = v
			}
		}
	}

	if prod == nil {
		return c.JSON(http.StatusOK, "not found")
	}
	return c.JSON(http.StatusOK, prod)
}

func postMas(c echo.Context) error {
	// e.Validator = &Valid{Valid: vNew}

	if err := c.Bind(&body); err != nil {
		return err
	}

	// if err := c.Validate(body); err != nil {
	// 	return err
	// }

	prod := map[int]string{
		len(mas) + 1: body.Name,
	}
	mas = append(mas, prod)
	return c.JSON(http.StatusOK, prod)
}

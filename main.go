package main

import (
	h "net/http"

	"github.com/labstack/echo/v4"

)

var (
	object = []map[int]string{{1: "one"}, {2: "two"}}
	e      = echo.New()
)

func main() {
	//start
	e.GET("/", func(c echo.Context) error {
		return c.JSON(h.StatusOK, object)
	})
	//end
	e.Logger.Fatal(e.Start(":3000"))
}

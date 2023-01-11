package start

import (
	h "net/http"

	"github.com/labstack/echo/v4"

)


func Start() {
	//start
	e.GET("/", func(c echo.Context) error {
		return c.JSON(h.StatusOK, object)
	})
	e.GET("/:id", getID)
	e.POST("/", addObject)
	e.PUT("/:id", putObject)
	e.DELETE("/:id", delObject)
	//end
	e.Logger.Fatal(e.Start(":3000"))

}

package start

import (
	"log"
	h "net/http"

	"github.com/labstack/echo/v4"

)

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("server MESSAGE")
		return next(c)
	}
}

func Start() {
	//start
	e.GET("/", func(c echo.Context) error {
		return c.JSON(h.StatusOK, object)
	})
	e.GET("/:id", getID, serverMessage)
	e.POST("/", addObject)
	e.PUT("/:id", putObject)
	e.DELETE("/:id", delObject)
	//end
	e.Logger.Fatal(e.Start(":3000"))

}

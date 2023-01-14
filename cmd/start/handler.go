package start

import (
	"log"
	h "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)


func serMes(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("you")

		return next(c)
	}
}

func Start() {
	//start
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(h.StatusOK, object)
	})
	e.GET("/api/:id", getID, serMes)
	e.POST("/api", addObject)
	e.PUT("/api/:id", putObject)
	e.DELETE("/api/:id", delObject)
	//end
	e.Logger.Fatal(e.Start(":3000"))

}

package start

import (
	"log"
	h "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("server MESSAGE")
		// c.Request().URL.Path = "/api/1"
		// fmt.Printf("%+v\n", c.Request())
		return next(c)
	}
}

func Start() {
	//start
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/api", func(c echo.Context) error {
		// fmt.Printf("inside = %+v\n", c.Request())
		return c.JSON(h.StatusOK, object)
	})
	e.GET("/api/:id", getID)
	e.POST("/api", addObject)
	e.PUT("/api/:id", putObject)
	e.DELETE("/api/:id", delObject, middleware.BodyLimit("1"))
	//end
	e.Logger.Fatal(e.Start(":3000"))

}

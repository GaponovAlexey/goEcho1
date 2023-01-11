package start

import (
	"go/echo/config"
	h "net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"

)

func init() {
	err := cleanenv.ReadEnv(&config.Cfg)
	
	cancel(err)

}

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

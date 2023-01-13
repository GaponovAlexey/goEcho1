package start

import (
	h "net/http"
	"strconv"

	"github.com/labstack/echo/v4"

)

var (
	object = []map[int]string{{1: "one"}, {2: "two"}}
	e      = echo.New()
)


// two
func getID(c echo.Context) error {
	gId, _ := strconv.Atoi(c.Param("id"))
	var getIdObject map[int]string
	for _, v := range object {
		for k := range v {
			if gId == k {
				getIdObject = v
			}
		}
	}
	// if getIdObject == nil {
	// 	return c.JSON(h.StatusNotFound, "not found")
	// }
	return c.JSON(h.StatusOK, getIdObject)
}

func addObject(c echo.Context) error {
	type res struct {
		Name string `json:"name"`
	}
	var body res

	if err := c.Bind(&body); err != nil {
		return err
	}

	newObject := map[int]string{
		len(object) + 1: body.Name,
	}

	object = append(object, newObject)

	return c.JSON(h.StatusOK, body)
}

func putObject(c echo.Context) error {
	gID, _ := strconv.Atoi(c.Param("id"))
	var newObject map[int]string
	for _, v := range object {
		for k := range v {
			if gID == k {
				newObject = v
			}
		}
	}

	type res struct {
		Name string `json:"name"`
	}
	var body res

	if err := c.Bind(&body); err != nil {
		return err
	}

	newObject[gID] = body.Name

	return c.JSON(h.StatusOK, newObject)
}

func delObject(c echo.Context) error {
	gId, _ := strconv.Atoi(c.Param("id"))

	var index int
	for i, v := range object {
		for k := range v {
			if gId == k {
				index = i
			}
		}
	}

	split := func(s []map[int]string, i int) []map[int]string {
		return append(s[:i], s[i+1:]...)
	}

	

	object = split(object, index)
	

	return c.JSON(h.StatusOK, "success")

}

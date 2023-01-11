package start

import (
	"fmt"
	h "net/http"
	"strconv"

	"github.com/labstack/echo/v4"

)

var (
	object = []map[int]string{{1: "one"}, {2: "two"}}
	e      = echo.New()
)

func cancel(e error) {
	if e != nil {
		fmt.Errorf("Error ->", e)
	}
}

func getID(c echo.Context) error {
	gId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// return fmt.Errorf("Err ->>>>", err)
	// }
	cancel(err)
	oneObject := map[int]string{}
	for _, v := range object {
		for k := range v {
			if k == gId {
				oneObject = v
			}
		}
	}
	if oneObject == nil {
		return c.JSON(h.StatusNotFound, "not found")
	}

	return c.JSON(h.StatusOK, oneObject)
}

func addObject(c echo.Context) error {
	type res struct {
		Name string `json:"name"`
	}
	var body res

	if err := c.Bind(&body); err != nil {
		return err
	}
	newObj := map[int]string{
		len(object) + 1: body.Name,
	}
	object = append(object, newObj)
	return c.JSON(h.StatusOK, newObj)
}

func putObject(c echo.Context) error {
	pId, err := strconv.Atoi(c.Param("id"))
	cancel(err)
	var newOb map[int]string
	for _, v := range object {
		for k := range v {
			if pId == k {
				newOb = v
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

	newOb[pId] = body.Name
	return c.JSON(h.StatusOK, newOb)

}

func delObject(c echo.Context) error {
	// var delOb map[int]string
	pID, err := strconv.Atoi(c.Param("id"))
	cancel(err)
	
	var index int
	for i, v := range object {
		for k := range v {
			if pID == k {
				// delOb = v
				index = i
			}
		}
	}

	splice := func(s []map[int]string, i int) []map[int]string {
		return append(s[:i], s[i+1:]...)
	}

	object = splice(object, index)
	return c.JSON(h.StatusOK, "success delete")

}

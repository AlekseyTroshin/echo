package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//example http://localhost:1323/show?team=x-men&member=wolverine

func main() {
	e := echo.New()
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+" member:"+member)
}

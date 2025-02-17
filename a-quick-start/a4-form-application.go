package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// example curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/saveApplication

func main() {
	e := echo.New()
	e.POST("/save", saveApplication)
	e.Logger.Fatal(e.Start(":1323"))
}

func saveApplication(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

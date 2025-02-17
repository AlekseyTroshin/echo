package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

//example curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/saveMultipart

func main() {
	e := echo.New()
	e.POST("/save", saveMultipart)
	e.Logger.Fatal(e.Start(":1323"))
}

func saveMultipart(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

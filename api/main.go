package main

import (
	"fmt"
	"net/http"

	"github.com/dylanwe/yifu/config"
	"github.com/dylanwe/yifu/database"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()
	c := config.GetConfig()
	database.Init()

	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.Logger.Fatal(server.Start(":8080"))
	fmt.Println("Started in " + string(c.Mode) + " mode")
}

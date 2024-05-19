package main

import (
	"fmt"

	"github.com/dylanwe/yifu/config"
	"github.com/dylanwe/yifu/database"
	"github.com/dylanwe/yifu/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()
	c := config.GetConfig()
	database.Init()

	server := echo.New()
	api := server.Group("/api")
	clothes := api.Group("/v1/clothes")
	routes.ClothesRoutes(clothes)

	fmt.Println("Starting in " + string(c.Mode) + " mode")
	server.Logger.Fatal(server.Start(":8080"))
}

package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dylanwe/yifu/config"
	"github.com/dylanwe/yifu/database"
	"github.com/dylanwe/yifu/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func main() {
	config.Init()
	c := config.GetConfig()
	database.Init()

	server := echo.New()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Middleware
	server.Use(slogecho.New(logger))
	server.Use(middleware.Recover())

	api := server.Group("/api")
	clothes := api.Group("/v1/clothes")
	routes.ClothesRoutes(clothes)
	outfits := api.Group("/v1/outfits")
	routes.OutfitsRoutes(outfits)

	fmt.Println("Starting in " + string(c.Mode) + " mode")
	server.Logger.Fatal(server.Start(":8080"))
}

package routes

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dylanwe/yifu/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ClothesRoutes(group *echo.Group) {
	group.POST("", func(c echo.Context) error {
		name := c.FormValue("name")
		color := c.FormValue("color")
		category := c.FormValue("category")
		image, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, Msg{"Image is required"})
		}

		extension := filepath.Ext(image.Filename)
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
			return c.JSON(http.StatusBadRequest, Msg{"Image must be a jpg, jpeg, or png"})
		}

		imageUrl, err := image.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Msg{"Error saving image"})
		}
		defer imageUrl.Close()

		dst, err := os.Create("images/" + image.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Msg{"Error saving image"})
		}
		defer dst.Close()

		if _, err = io.Copy(dst, imageUrl); err != nil {
			return c.JSON(http.StatusInternalServerError, Msg{"Error saving image"})
		}

		f, _ := os.Getwd()
		imagePath := f + "/images/" + image.Filename

		clothing := database.Clothing{
			Id:       uuid.New(),
			Name:     name,
			Color:    color,
			Category: category,
			Image:    imagePath,
		}

		database.DB.Create(&clothing)

		return c.JSON(http.StatusCreated, Msg{"Clothing created successfully!"})
	})

	group.GET("", func(c echo.Context) error {
		var clothes []database.Clothing
		database.DB.Find(&clothes)
		return c.JSON(http.StatusOK, clothes)
	})

	group.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		var clothing database.Clothing
		database.DB.Where("id = ?", id).First(&clothing)

		if clothing.Id == uuid.Nil {
			return c.JSON(http.StatusNotFound, Msg{"Clothing not found"})
		}

		return c.JSON(http.StatusOK, clothing)
	})

	group.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")
		var clothing database.Clothing
		database.DB.Where("id = ?", id).First(&clothing)

		if clothing.Id == uuid.Nil {
			return c.JSON(http.StatusNotFound, Msg{"Clothing not found"})
		}

		database.DB.Delete(&clothing)
		return c.JSON(http.StatusOK, Msg{"Clothing deleted successfully!"})
	})
}

package routes

import (
	"net/http"

	"github.com/dylanwe/yifu/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func OutfitsRoutes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		var outfits []database.Outfit
		database.DB.Find(&outfits)

		for i := range outfits {
			database.DB.Model(&outfits[i]).Association("Clothing").Find(&outfits[i].Clothing)
		}

		return c.JSON(http.StatusOK, outfits)
	})

	group.POST("", func(c echo.Context) error {
		var outfitRequest OutfitRequest
		if err := c.Bind(&outfitRequest); err != nil {
			return c.JSON(http.StatusBadRequest, Msg{"Invalid outfit request"})
		}

		outfit := database.Outfit{
			Name:     outfitRequest.Name,
			Clothing: []database.Clothing{},
		}

		database.DB.Create(&outfit)
		return c.JSON(http.StatusCreated, Msg{"Outfit created successfully!"})
	})

	group.GET("/:id", func(c echo.Context) error {
		var outfit database.Outfit
		id := c.Param("id")
		database.DB.First(&outfit, "id = ?", id)

		if outfit.Id == uuid.Nil {
			return c.JSON(http.StatusNotFound, Msg{"Outfit not found"})
		}

		database.DB.Model(&outfit).Association("Clothing").Find(&outfit.Clothing)
		return c.JSON(http.StatusOK, outfit)
	})

	group.DELETE("/:id", func(c echo.Context) error {
		var outfit database.Outfit
		id := c.Param("id")
		database.DB.First(&outfit, "id = ?", id)

		if outfit.Id == uuid.Nil {
			return c.JSON(http.StatusNotFound, Msg{"Outfit not found"})
		}

		database.DB.Delete(&outfit)
		return c.JSON(http.StatusOK, "Outfit deleted successfully!")
	})

	group.POST("/:id/clothes", func(c echo.Context) error {
		var outfitClothingRequest OutfitClothingRequest
		if err := c.Bind(&outfitClothingRequest); err != nil {
			return c.JSON(http.StatusBadRequest, Msg{"Invalid outfit clothing request"})
		}

		var outfit database.Outfit
		id := c.Param("id")
		database.DB.First(&outfit, "id = ?", id)

		if outfit.Id == uuid.Nil {
			return c.JSON(http.StatusNotFound, Msg{"Outfit not found"})
		}

		var clothingIds []uuid.UUID
		for _, clothingId := range outfitClothingRequest.ClothingIds {
			clothingIds = append(clothingIds, uuid.MustParse(clothingId))
		}

		var clothing []database.Clothing
		database.DB.Find(&clothing, clothingIds)

		database.DB.Model(&outfit).Association("Clothing").Append(&clothing)
		return c.JSON(http.StatusOK, Msg{"Clothing added to outfit successfully!"})
	})

	group.DELETE("/:id/clothes", func(c echo.Context) error {
		var outfitClothingRequest OutfitClothingRequest
		if err := c.Bind(&outfitClothingRequest); err != nil {
			return c.JSON(http.StatusBadRequest, Msg{"Invalid outfit clothing request"})
		}

		var outfit database.Outfit
		id := c.Param("id")
		database.DB.First(&outfit, "id = ?", id)

		if outfit.Id == uuid.Nil {
			return c.JSON(http.StatusNotFound, Msg{"Outfit not found"})
		}

		var clothingIds []uuid.UUID
		for _, clothingId := range outfitClothingRequest.ClothingIds {
			clothingIds = append(clothingIds, uuid.MustParse(clothingId))
		}

		var clothing []database.Clothing
		database.DB.Find(&clothing, clothingIds)

		database.DB.Model(&outfit).Association("Clothing").Delete(&clothing)
		return c.JSON(http.StatusOK, Msg{"Clothing removed from outfit successfully!"})
	})
}

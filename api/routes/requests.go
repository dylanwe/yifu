package routes

type OutfitRequest struct {
	Name string `json:"name"`
}

type OutfitClothingRequest struct {
	ClothingIds []string `json:"clothingIds"`
}

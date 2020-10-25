package ingredient

// IngredientListItemResponseDTO godoc
type IngredientListItemResponseDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Quantity uint32 `json:"quantity"`
}

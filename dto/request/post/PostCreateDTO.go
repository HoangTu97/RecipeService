package PostRequest

type PostCreateRequestDTO struct {
	Description string   `json:"description"`
	HashTags    []string `json:"tags"`

	CateID uint `json:"cateId"`

	RecipeImage       string   `json:"recipeImage"`
	RecipeName        string   `json:"recipeName"`
	RecipeDuration    uint32   `json:"recipeDuration"`
	RecipeIngredients []string `json:"recipeIngredients"`
	RecipeSteps       []string `json:"recipeSteps"`
}

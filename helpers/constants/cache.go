package constants

var CACHE cache

type cache struct {
	ARTILE string
	TAG string
	USER string
	IMAGE string
	RECIPE string
	CATEGORY string
	COMMENT string
	POST string
	INGREDIENT string
}

func init() {
	CACHE = cache{
		ARTILE: "ARTILE",
		TAG: "TAG",
		USER: "USER",
		IMAGE: "IMAGE",
		RECIPE: "RECIPE",
		CATEGORY: "CATEGORY",
		COMMENT: "COMMENT",
		POST: "POST",
		INGREDIENT: "INGREDIENT",
	}
}

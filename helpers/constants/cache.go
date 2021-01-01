package constants

var CACHE cache

type cache struct {
	ARTILE string
	TAG string
}

func init() {
	CACHE = cache{
		ARTILE: "ARTILE",
		TAG: "TAG",
	}
}

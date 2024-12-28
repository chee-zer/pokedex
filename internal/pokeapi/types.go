package pokeapi

type LocationRes struct {
	Name string
	Url string
}

type ApiRes struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LocationRes `json:"results"`
}
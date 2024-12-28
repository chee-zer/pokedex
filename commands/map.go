package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chee-zer/pokedex/internal/pokeapi"
)

func printLocations(locations []pokeapi.LocationRes) {
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
}

func Map() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("couldn't make request")
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("invalid request")
	}

	var jsonData pokeapi.ApiRes
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&jsonData); err != nil {
		return fmt.Errorf("could not decode response")
	}

	printLocations(jsonData.Results)
	 return nil
}
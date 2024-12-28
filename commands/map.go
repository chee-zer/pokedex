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

func Map(cfg *Config) error {
	if cfg.Next == "" {
		return fmt.Errorf("you're on the last page")
	}
	url := cfg.Next
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

	cfg.Next = jsonData.Next
	cfg.Previous = jsonData.Previous

	 return nil
}

func Mapb(cfg *Config) error {
	if cfg.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	url := cfg.Previous
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

	cfg.Next = jsonData.Next
	cfg.Previous = jsonData.Previous

	 return nil
}


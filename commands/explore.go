package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemons struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func Explore(cfg *Config) error {
	if len(cfg.Args) == 0 {
		fmt.Println("Please provide a location to explore!")
		return nil
	}
	location := cfg.Args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + location
	cachedData, exists := cfg.C.Get(url)
	if exists {
		printPokemons(cachedData, location)
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		fetchedData, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.C.Add(url, fetchedData)

		printPokemons(fetchedData, location)
	}
	return nil
}

func printPokemons(val []byte, location string) {
	var data pokemons
	if err := json.Unmarshal(val, &data); err != nil {
		fmt.Println("Not a valid location!")
		return
	}
	fmt.Printf("Exploring %v...\nFound Pokemon:\n", location)
	for _, val := range data.PokemonEncounters {
		fmt.Printf("- %v\n",val.Pokemon.Name)
	}
}

package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chee-zer/pokedex/internal/pokeapi"
)

func printLocations(locations []pokeapi.LocationRes) {
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
}



func Map(cfg *Config) error {
	//checks if there are no more pages
	if cfg.Next == "" {
		return fmt.Errorf("you're on the last page")
	}
	//if next page exists
	url := cfg.Next
	var jsonData pokeapi.ApiRes
	var data []byte

	//check if the next url is cached
	cachedData, exists := cfg.C.Get(url)
	if exists {
		data = cachedData
	} else {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("couldn't make request")
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("invalid request")
	}
	
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not process fetched data")
	}
	

	cfg.C.Add(url, data)
	}
	if err := json.Unmarshal(data, &jsonData); err != nil {
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
	var jsonData pokeapi.ApiRes
	var data []byte
	cachedData, exists := cfg.C.Get(url)
	if exists {
		data = cachedData
		} else {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return fmt.Errorf("couldn't make request")
			}
			
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("invalid request")
			}
			defer res.Body.Close()
			
			data, err = io.ReadAll(res.Body)
			if err != nil {
				return fmt.Errorf("could not process fetched data")
			}

	cfg.C.Add(url, data)
	}
	
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("could not decode response")
	}
	printLocations(jsonData.Results)

	cfg.Next = jsonData.Next
	cfg.Previous = jsonData.Previous

	 return nil
}


package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
)

func Catch(cfg *Config) error {
	//arg, check in cache, catch, fetc, store in cache, catch
	mon := cfg.Args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", mon)
	_, exists := cfg.Pokedex.Entry[mon]
	if exists {
		fmt.Printf("%v is already in your collection\n", mon)
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + mon
	var data []byte
	cachedData, exists := cfg.C.Get(url)
	if exists {
		data = cachedData
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
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.C.Add(url, data)
	}
	var poke pokedexRaw
	if err := json.Unmarshal(data, &poke); err != nil {
		return err

	}
	caught := catchPokemon(poke.Basexp)
	if caught {
		storePoke := cookData(poke)

		cfg.Pokedex.Add(mon, storePoke)
		fmt.Printf("%v was caught!\n", mon)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%v escaped!\n", mon)
	}
	return nil

}

func catchPokemon(baseXP int) bool {
	chance := 100 - int(math.Floor(float64(baseXP)/10.0))
	ra := rand.Intn(200)
	fmt.Printf("baseXP: %v, chance: %v, ra: %v\n", baseXP, chance, ra)
	if ra <= chance {
		return true
	} else {
		return false
	}
}

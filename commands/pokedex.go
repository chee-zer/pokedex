package commands

import "fmt"

type pokedexRaw struct {
	Name   string `json:"name"`
	Basexp int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Typee struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Pokemon struct {
	Name   string
	Basexp int
	Height int
	Weight int
	Stats  map[string]int
	Types  []string
}

type Pokedex struct {
	Entry map[string]Pokemon
}

func cookData(dat pokedexRaw) Pokemon {
	var cookedData Pokemon
	cookedData.Stats = make(map[string]int)
	cookedData.Name = dat.Name
	cookedData.Basexp = dat.Basexp
	cookedData.Height = dat.Height
	cookedData.Weight = dat.Weight
	for _, stat := range dat.Stats {
		cookedData.Stats[stat.Stat.Name] = stat.BaseStat
	}
	for _, ty := range dat.Types {
		cookedData.Types = append(cookedData.Types, ty.Typee.Name)
	}

	return cookedData
}

func (p *Pokedex) Add(name string, val Pokemon) error {

	p.Entry[name] = val
	return nil
}

func PokedexC(cfg *Config) error {
	if len(cfg.Pokedex.Entry) == 0 {
		fmt.Printf("Your Pokedex is empty!\nCatch some pokemon to add them to your Pokedex!\n")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for k := range cfg.Pokedex.Entry {
		fmt.Printf("- %v\n", k)
	}
	return nil
}

package commands

import "fmt"

func Inspect(cfg *Config) error {
	mon := cfg.Args[0]
	poke, exists := cfg.Pokedex.Entry[mon]
	if exists {
		fmt.Printf("Name: %v\n", poke.Name)
		fmt.Printf("Height: %v\n", poke.Height)
		fmt.Printf("Weight: %v\n", poke.Weight)
		fmt.Println("Stats:")
		for stat, value := range poke.Stats {
			fmt.Printf("-%v: %v\n", stat, value)
		}
		fmt.Println("Types:")
		for _, v := range poke.Types {
			fmt.Printf("- %v\n", v)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

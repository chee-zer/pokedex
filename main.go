package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chee-zer/pokedex/commands"
	"github.com/chee-zer/pokedex/internal/pokecache"
)

func main() {
	//registry of valid commands
	commands.RegisterCommands()
	urlCache := pokecache.NewCache(5 * time.Second)

	cfg := commands.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		C:        urlCache,
	}

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if s.Scan() {
			if s.Text() == "" {
				continue
			}
			inp := cleanInput(s.Text())[0]
			cfg.Args = cleanInput(s.Text())[1:]
			cmd, exists := commands.CmdReg[inp]
			if exists {
				err := cmd.Callback(&cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

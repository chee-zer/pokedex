package commands

import "github.com/chee-zer/pokedex/internal/pokecache"

// cliCommand can stay unexported, but its properties need to
// be exported, since they are being used indirectly outside this package
type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
	C        *pokecache.Cache
	Pokedex  Pokedex
	Args     []string
}

var CmdReg = map[string]CliCommand{}

func RegisterCommands() {
	CmdReg["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    Exit,
	}
	CmdReg["help"] = CliCommand{
		Name:        "help",
		Description: "Display a help message",
		Callback:    Help,
	}
	CmdReg["map"] = CliCommand{
		Name:        "map",
		Description: "Display next 20 names of location areas in the Pokemon world.",
		Callback:    Map,
	}
	CmdReg["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Display the previous 20 names of location areas in the Pokemon world.",
		Callback:    Mapb,
	}
	CmdReg["explore"] = CliCommand{
		Name:        "explore",
		Description: "Explore the location for pokemons!",
		Callback:    Explore,
	}
	CmdReg["catch"] = CliCommand{
		Name:        "catch",
		Description: "Catch a Pokemon duh",
		Callback:    Catch,
	}
	CmdReg["inspect"] = CliCommand{
		Name:        "inspect",
		Description: "Inspect a pokemon",
		Callback:    Inspect,
	}
	CmdReg["pokedex"] = CliCommand{
		Name:        "pokedex",
		Description: "List all caught pokemon",
		Callback:    PokedexC,
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var cmdReg map[string]cliCommand

func main()  {
	//registry of valid commands
	cmdReg = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
	}
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if s.Scan() {
			inp := cleanInput(s.Text())[0];
			cmd, exists := cmdReg[inp]
			if exists {
				cmd.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

//callback for commands

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for cmd, val := range cmdReg {
		fmt.Printf("%v: %v\n", cmd, val.description)
	}
	return nil
}
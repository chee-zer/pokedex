package commands

import "fmt"

func getCmdReg () map[string]CliCommand {
	return CmdReg
}

func Help() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	cmdReg := getCmdReg()
	for cmd, val := range cmdReg {
		fmt.Printf("%v: %v\n\n", cmd, val.Description)
	}
	return nil
}
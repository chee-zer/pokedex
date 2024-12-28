package commands

//cliCommand can stay unexported, but its properties need to
//be exported, since they are being used indirectly outside this package
type CliCommand struct {
	Name string
	Description string
	Callback func() error
}



var CmdReg = map[string]CliCommand{}

func RegisterCommands() {
	CmdReg["exit"] = CliCommand {
		Name: "exit",
		Description: "Exit the Pokedex",
		Callback: Exit,
	}
	CmdReg["help"] = CliCommand{
		Name: "help",
		Description: "Display a help message",
		Callback: Help,
	}
	CmdReg["map"] = CliCommand{
		Name: "map",
		Description: "Display names of location areas in the Pokemon world.",
		Callback: Map,
	}
}
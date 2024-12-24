package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	for {
		fmt.Print("Pokedex > ")
		s := bufio.NewScanner(os.Stdin)
		if s.Scan() {
			op := cleanInput(s.Text())
			fmt.Printf("Your command was: %v\n", op[0])
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
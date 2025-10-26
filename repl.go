package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Structs
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Functions
func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Closing")
}

func commandHelp() error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return errors.New("Help")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}

// End of Functions

func startRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Read error: ", err)
			}
		}

		userInput := scanner.Text()
		processedInput := cleanInput(userInput)

		command := processedInput[0]
		if value, ok := commands[command]; ok {
			value.callback()
		}
	}
}

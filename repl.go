package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cantr1/GoPokedex/internal/pokeapi"
)

// Structs
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

// Functions
func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Closing")
}

func commandHelp(cfg *config) error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return errors.New("Help")
}

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
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
		"map": {
			name:        "map",
			description: "Displays a list of all locations",
			callback:    commandMap,
		},
	}
}

// End of Functions

func startRepl(cfg *config) {
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
			value.callback(cfg)
		}
	}
}

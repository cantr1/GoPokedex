package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		processedInput := cleanInput(userInput)
		fmt.Printf("Your command was: %v\n", processedInput[0])

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input: %v\n", err)
		}
	}
}

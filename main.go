package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}

func main() {
	myString := "Charmander Bulbasaur PIKACHU"
	fmt.Println(myString)
	fmt.Println(cleanInput(myString))
}

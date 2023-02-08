//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
)

// the questions to ask
var defaultPasswordCharacterPrompt = &inquirer.Password{
	Message: "What is your password? (Default hide character)",
}
var customPasswordCharacterPrompt = &inquirer.Password{
	Message: "What is your password? (Custom hide character)",
}

func main() {

	var defaultPass string
	var customPass string

	// ask the question
	err := inquirer.AskOne(defaultPasswordCharacterPrompt, &defaultPass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println()
	err = inquirer.AskOne(customPasswordCharacterPrompt, &customPass, inquirer.WithHideCharacter('-'))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("Password 1: %s.\n Password 2: %s\n", defaultPass, customPass)
}

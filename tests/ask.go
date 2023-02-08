//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
)

// the questions to ask
var simpleQs = []*inquirer.Question{
	{
		Name: "name",
		Prompt: &inquirer.Input{
			Message: "What is your name?",
			Default: "Johnny Appleseed",
		},
	},
	{
		Name: "color",
		Prompt: &inquirer.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green", "yellow"},
			Default: "yellow",
		},
		Validate: inquirer.Required,
	},
}

var singlePrompt = &inquirer.Input{
	Message: "What is your name?",
	Default: "Johnny Appleseed",
}

func main() {

	fmt.Println("Asking many.")
	// a place to store the answers
	ans := struct {
		Name  string
		Color string
	}{}
	err := inquirer.Ask(simpleQs, &ans)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Asking one.")
	answer := ""

	err = inquirer.AskOne(singlePrompt, &answer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Answered with %v.\n", answer)

	fmt.Println("Asking one with validation.")
	vAns := ""
	err = inquirer.AskOne(&inquirer.Input{Message: "What is your name?"}, &vAns, inquirer.WithValidator(inquirer.Required))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Answered with %v.\n", vAns)
}

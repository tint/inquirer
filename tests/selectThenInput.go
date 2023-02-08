//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
)

// the questions to ask
var simpleQs = []*inquirer.Question{
	{
		Name: "color",
		Prompt: &inquirer.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: inquirer.Required,
	},
	{
		Name: "name",
		Prompt: &inquirer.Input{
			Message: "What is your name?",
		},
		Validate: inquirer.Required,
	},
}

func main() {
	answers := struct {
		Color string
		Name  string
	}{}
	// ask the question
	err := inquirer.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("%s chose %s.\n", answers.Name, answers.Color)
}
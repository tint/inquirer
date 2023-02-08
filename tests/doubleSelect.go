//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
)

var simpleQs = []*inquirer.Question{
	{
		Name: "color",
		Prompt: &inquirer.Select{
			Message: "select1:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: inquirer.Required,
	},
	{
		Name: "color2",
		Prompt: &inquirer.Select{
			Message: "select2:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: inquirer.Required,
	},
}

func main() {
	answers := struct {
		Color  string
		Color2 string
	}{}
	// ask the question
	err := inquirer.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("%s and %s.\n", answers.Color, answers.Color2)
}

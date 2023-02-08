//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
)

// the questions to ask
var simpleQs = []*inquirer.Question{
	{
		Name: "letter",
		Prompt: &inquirer.Select{
			Message: "Choose a letter:",
			Options: []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
				"g",
				"h",
				"i",
				"j",
			},
		},
		Validate: inquirer.Required,
	},
}

func main() {
	answers := struct {
		Letter string
	}{}

	// ask the question
	err := inquirer.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("you chose %s.\n", answers.Letter)
}

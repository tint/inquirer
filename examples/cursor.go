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
		},
		Validate: inquirer.Required,
	},
}

func main() {
	ansmap := make(map[string]interface{})

	// ask the question
	err := inquirer.Ask(simpleQs, &ansmap, inquirer.WithShowCursor(true))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("Your name is %s.\n", ansmap["name"])
}

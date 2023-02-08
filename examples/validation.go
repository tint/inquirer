//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
)

// the questions to ask
var validationQs = []*inquirer.Question{
	{
		Name:     "name",
		Prompt:   &inquirer.Input{Message: "What is your name?"},
		Validate: inquirer.Required,
	},
	{
		Name:   "valid",
		Prompt: &inquirer.Input{Message: "Enter 'foo':", Default: "not foo"},
		Validate: func(val interface{}) error {
			// if the input matches the expectation
			if str := val.(string); str != "foo" {
				return fmt.Errorf("You entered %s, not 'foo'.", str)
			}
			// nothing was wrong
			return nil
		},
	},
}

func main() {
	// the place to hold the answers
	answers := struct {
		Name  string
		Valid string
	}{}
	err := inquirer.Ask(validationQs, &answers)

	if err != nil {
		fmt.Println("\n", err.Error())
	}
}

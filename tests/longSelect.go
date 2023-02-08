//go:build ignore

package main

import "github.com/tint/inquirer"

func main() {
	color := ""
	prompt := &inquirer.Select{
		Message: "Choose a color:",
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
	}
	inquirer.AskOne(prompt, &color)
}

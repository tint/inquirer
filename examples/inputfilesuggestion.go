//go:build ignore

package main

import (
	"fmt"
	"path/filepath"

	"github.com/tint/inquirer"
)

func suggestFiles(toComplete string) []string {
	files, _ := filepath.Glob(toComplete + "*")
	return files
}

// the questions to ask
var q = []*inquirer.Question{
	{
		Name: "file",
		Prompt: &inquirer.Input{
			Message: "Which file should be read?",
			Suggest: suggestFiles,
			Help:    "Any file; do not need to exist yet",
		},
		Validate: inquirer.Required,
	},
}

func main() {
	answers := struct {
		File string
	}{}

	// ask the question
	err := inquirer.Ask(q, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("File chosen %s.\n", answers.File)
}

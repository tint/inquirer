//go:build ignore

package main

import (
	"fmt"
	"strings"

	"github.com/tint/inquirer"
	TestUtil "github.com/tint/inquirer/tests/util"
)

var answer = ""

var goodTable = []TestUtil.TestTableEntry{
	{
		"should open in editor", &inquirer.Editor{
			Message: "should open",
		}, &answer, nil,
	},
	{
		"has help", &inquirer.Editor{
			Message: "press ? to see message",
			Help:    "Does this work?",
		}, &answer, nil,
	},
	{
		"should not include the default value in the prompt", &inquirer.Editor{
			Message:     "the default value 'Hello World' should not include in the prompt",
			HideDefault: true,
			Default:     "Hello World",
		}, &answer, nil,
	},
	{
		"should write the default value to the temporary file before the launch of the editor", &inquirer.Editor{
			Message:       "the default value 'Hello World' is written to the temporary file before the launch of the editor",
			AppendDefault: true,
			Default:       "Hello World",
		}, &answer, nil,
	},
	{
		Name: "should print the validation error, and recall the submitted invalid value instead of the default",
		Prompt: &inquirer.Editor{
			Message:       "the default value 'Hello World' is written to the temporary file before the launch of the editor",
			AppendDefault: true,
			Default:       `this is the default value. change it to something containing "invalid" (in vi type "ccinvalid<Esc>ZZ")`,
		},
		Value: &answer,
		Validate: func(v interface{}) error {
			s := v.(string)
			if strings.Contains(s, "invalid") {
				return fmt.Errorf(`this is the error message. change the input to something not containing "invalid"`)
			}
			return nil
		},
	},
}

func main() {
	TestUtil.RunTable(goodTable)
}

//go:build ignore

package main

import (
	"fmt"

	"github.com/tint/inquirer"
	TestUtil "github.com/tint/inquirer/tests/util"
)

var answer = []string{}

var table = []TestUtil.TestTableEntry{
	{
		"standard", &inquirer.MultiSelect{
			Message: "What days do you prefer:",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		}, &answer, nil,
	},
	{
		"default (sunday, tuesday)", &inquirer.MultiSelect{
			Message: "What days do you prefer:",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Default: []string{"Sunday", "Tuesday"},
		}, &answer, nil,
	},
	{
		"default not found", &inquirer.MultiSelect{
			Message: "What days do you prefer:",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Default: []string{"Sundayaa"},
		}, &answer, nil,
	},
	{
		"no help - type ?", &inquirer.MultiSelect{
			Message: "What days do you prefer:",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Default: []string{"Sundayaa"},
		}, &answer, nil,
	},
	{
		"can navigate with j/k", &inquirer.MultiSelect{
			Message: "What days do you prefer:",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Default: []string{"Sundayaa"},
		}, &answer, nil,
	},
	{
		"descriptions", &inquirer.MultiSelect{
			Message: "What days do you prefer:",
			Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			Description: func(value string, index int) string {
				return value + fmt.Sprint(index)

			},
		}, &answer, nil,
	},
}

func main() {
	TestUtil.RunTable(table)
}

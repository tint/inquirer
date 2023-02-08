//go:build ignore

package main

import (
	"github.com/tint/inquirer"
	TestUtil "github.com/tint/inquirer/tests/util"
)

var answer = false

var goodTable = []TestUtil.TestTableEntry{
	{
		"Enter 'yes'", &inquirer.Confirm{
			Message: "yes:",
		}, &answer, nil,
	},
	{
		"Enter 'no'", &inquirer.Confirm{
			Message: "yes:",
		}, &answer, nil,
	},
	{
		"default", &inquirer.Confirm{
			Message: "yes:",
			Default: true,
		}, &answer, nil,
	},
	{
		"not recognized (enter random letter)", &inquirer.Confirm{
			Message: "yes:",
			Default: true,
		}, &answer, nil,
	},
	{
		"no help - type '?'", &inquirer.Confirm{
			Message: "yes:",
			Default: true,
		}, &answer, nil,
	},
}

func main() {
	TestUtil.RunTable(goodTable)
}

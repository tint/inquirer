//go:build ignore

package main

import (
	"github.com/tint/inquirer"
	TestUtil "github.com/tint/inquirer/tests/util"
)

var value = ""

var table = []TestUtil.TestTableEntry{
	{
		"standard", &inquirer.Password{Message: "Please type your password:"}, &value, nil,
	},
	{
		"please make sure paste works", &inquirer.Password{Message: "Please paste your password:"}, &value, nil,
	},
	{
		"no help, send '?'", &inquirer.Password{Message: "Please type your password:"}, &value, nil,
	},
}

func main() {
	TestUtil.RunTable(table)
}

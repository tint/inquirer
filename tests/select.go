//go:build ignore

package main

import (
	"github.com/tint/inquirer"
	TestUtil "github.com/tint/inquirer/tests/util"
)

var answer = ""

var goodTable = []TestUtil.TestTableEntry{
	{
		"standard", &inquirer.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
		}, &answer, nil,
	},
	{
		"short", &inquirer.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue"},
		}, &answer, nil,
	},
	{
		"default", &inquirer.Select{
			Message: "Choose a color (should default blue):",
			Options: []string{"red", "blue", "green"},
			Default: "blue",
		}, &answer, nil,
	},
	{
		"one", &inquirer.Select{
			Message: "Choose one:",
			Options: []string{"hello"},
		}, &answer, nil,
	},
	{
		"no help, type ?", &inquirer.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue"},
		}, &answer, nil,
	},
	{
		"passes through bottom", &inquirer.Select{
			Message: "Choose one:",
			Options: []string{"red", "blue"},
		}, &answer, nil,
	},
	{
		"passes through top", &inquirer.Select{
			Message: "Choose one:",
			Options: []string{"red", "blue"},
		}, &answer, nil,
	},
	{
		"can navigate with j/k", &inquirer.Select{
			Message: "Choose one:",
			Options: []string{"red", "blue", "green"},
		}, &answer, nil,
	},
}

var badTable = []TestUtil.TestTableEntry{
	{
		"no options", &inquirer.Select{
			Message: "Choose one:",
		}, &answer, nil,
	},
}

func main() {
	TestUtil.RunTable(goodTable)
	TestUtil.RunErrorTable(badTable)
}

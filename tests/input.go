//go:build ignore

package main

import (
	"github.com/tint/inquirer"
	TestUtil "github.com/tint/inquirer/tests/util"
)

var val = ""

var table = []TestUtil.TestTableEntry{
	{
		"no default", &inquirer.Input{Message: "Hello world"}, &val, nil,
	},
	{
		"default", &inquirer.Input{Message: "Hello world", Default: "default"}, &val, nil,
	},
	{
		"no help, send '?'", &inquirer.Input{Message: "Hello world"}, &val, nil,
	},
	{
		"Home, End Button test in random location", &inquirer.Input{Message: "Hello world"}, &val, nil,
	}, {
		"Delete and forward delete test at random location (test if screen overflows)", &inquirer.Input{Message: "Hello world"}, &val, nil,
	}, {
		"Moving around lines with left & right arrow keys", &inquirer.Input{Message: "Hello world"}, &val, nil,
	}, {
		"Runes with width > 1. Enter 一 you get to the next line", &inquirer.Input{Message: "Hello world"}, &val, nil,
	},
}

func main() {
	TestUtil.RunTable(table)
}

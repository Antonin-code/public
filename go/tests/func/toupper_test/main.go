package main

import (
	"strings"

	student "student"

	"lib"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello! How are you?")
	for _, arg := range table {
		lib.Challenge("ToUpper", student.ToUpper, strings.ToUpper, arg)
	}
}

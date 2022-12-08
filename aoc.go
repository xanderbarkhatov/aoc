package aoc

import (
	"os"
	"strings"
)

func ReadInput(prod bool) string {
	file := "input.dev"

	if prod {
		file = "input.prod"
	}

	input, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return string(input)
}

func ReadLines(prod bool) []string {
	return strings.Split(ReadInput(prod), "\n")
}

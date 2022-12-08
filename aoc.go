package aoc

import "os"

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

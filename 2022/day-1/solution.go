package main

import (
	"aoc"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := aoc.ReadInput(true)

	elves := strings.Split(input, "\n\n")
	calories := make([]int, len(elves))

	for i, elf := range elves {
		for _, item := range strings.Split(elf, "\n") {
			n, e := strconv.Atoi(item)

			if e != nil {
				panic(e)
			}

			calories[i] += n
		}

	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	partOne := calories[0]
	partTwo := calories[0] + calories[1] + calories[2]

	println(partOne)
	println(partTwo)

}

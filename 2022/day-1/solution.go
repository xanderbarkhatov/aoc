package main

import (
	"aoc"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := aoc.ReadInput(true)

	elves := strings.Split(input, "\n\n")
	calories := make([]int, len(elves))

	for i, elf := range elves {
		items := strings.Split(elf, "\n")
		sum := 0

		for _, item := range items {
			n, e := strconv.Atoi(item)

			if e != nil {
				panic(e)
			}

			sum += n
		}

		calories[i] = sum
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	partOne := calories[0]
	partTwo := 0

	for _, v := range calories[:3] {
		partTwo += v
	}

	fmt.Printf("%d\n", partOne)
	fmt.Printf("%d\n", partTwo)
}

package main

import (
	"aoc"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func sumCalories(elves []string) []int {
	calories := make([]int, len(elves))

	for i, elf := range elves {
		items := strings.Split(elf, "\n")
		sum := 0

		for _, item := range items {
			n, e := strconv.Atoi(item)
			handleError(e)
			sum += n
		}

		calories[i] = sum
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return calories
}

func partOne(input string) int {
	elves := strings.Split(input, "\n\n")
	calories := sumCalories(elves)

	return calories[0]
}

func partTwo(input string) int {
	elves := strings.Split(input, "\n\n")
	calories := sumCalories(elves)
	total := 0

	for _, v := range calories[:3] {
		total += v
	}

	return total
}

func main() {
	input := aoc.ReadInput(true)

	fmt.Printf("%d\n", partOne(input))
	fmt.Printf("%d\n", partTwo(input))
}

package main

import (
	"aoc"
	"strings"
)

func main() {
	lines := aoc.ReadLines(true)

	shapeToScore := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	outcomes := []int{3, 0, 6}
	score := 0

	for _, line := range lines {
		shapes := strings.Split(line, " ")
		oScore := shapeToScore[shapes[0]]
		pScore := shapeToScore[shapes[len(shapes)-1]]
		score += pScore + outcomes[(oScore-pScore+3)%3]
	}

	println(score)
}

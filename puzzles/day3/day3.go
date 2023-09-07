package puzzles

import (
	"strings"

	"github.com/anshulkanwar/aoc-go/internal"
)

func parseInput() [][]string {
	inputStr := internal.ReadInputFile("puzzles/day3/input.txt")

	var input [][]string

	for _, contents := range strings.Split(inputStr, "\n") {
		nItems := len(contents)

		input = append(input, []string{contents[:nItems/2], contents[nItems/2:]})
	}

	return input
}

func getPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 97 + 1
	} else {
		return int(item) - 65 + 27
	}
}

func SolveDay3Part1() {
	input := parseInput()

	answer := 0

	for _, contents := range input {
		firstHalf, secondHalf := contents[0], contents[1]

	outerLoop:
		for _, itemInFirstHalf := range firstHalf {
			for _, itemInSecondHalf := range secondHalf {
				if itemInFirstHalf == itemInSecondHalf {
					answer += getPriority(itemInFirstHalf)
					break outerLoop
				}
			}
		}
	}

	println(answer)
}

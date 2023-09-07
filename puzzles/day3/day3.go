package puzzles

import (
	"fmt"
	"strings"

	"github.com/anshulkanwar/aoc-go/internal"
)

func readInput() []string {
	inputStr := internal.ReadInputFile("puzzles/day3/input.txt")

	return strings.Split(inputStr, "\n")
}

func parseInput() [][]string {
	var input [][]string

	for _, contents := range readInput() {
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

	fmt.Println(answer)
}

func groupOfThree(input []string) [][]string {
	var output [][]string

	idx := 0

	for idx < len(input) {
		var group []string
		group = append(group, input[idx:idx+3]...)

		output = append(output, group)
		idx += 3
	}

	return output
}

func SolveDay3Part2() {
	input := readInput()
	groups := groupOfThree(input)
	answer := 0

	for _, group := range groups {
	forEachGroup:
		for _, elf1Items := range group[0] {
			for _, elf2Items := range group[1] {
				for _, elf3Items := range group[2] {
					if elf1Items == elf2Items && elf1Items == elf3Items {
						answer += getPriority((elf1Items))
						break forEachGroup
					}
				}
			}
		}
	}

	fmt.Println(answer)
}

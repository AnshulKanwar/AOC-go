package puzzles

import (
	"strconv"
	"strings"

	"github.com/anshulkanwar/aoc-go/internal"
)

func ParseInput() [][][]int {
	inputStr := internal.ReadInputFile("puzzles/day4/input.txt")
	var input [][][]int

	for _, sectionPairStr := range strings.Split(inputStr, "\n") {
		var sectionPair [][]int

		for _, sectionStr := range strings.Split(sectionPairStr, ",") {
			section := strings.Split(sectionStr, "-")
			first, _ := strconv.Atoi(section[0])
			second, _ := strconv.Atoi(section[1])
			sectionPair = append(sectionPair, []int{first, second})
		}

		input = append(input, sectionPair)
	}

	return input
}

func SovleDay4Part1(input [][][]int) int {
	answer := 0

	for _, sectionPair := range input {
		elf1First, elf1Second := sectionPair[0][0], sectionPair[0][1]
		elf2First, elf2Second := sectionPair[1][0], sectionPair[1][1]

		if elf2First >= elf1First && elf2Second <= elf1Second {
			answer++
		} else if elf1First >= elf2First && elf1Second <= elf2Second {
			answer++
		}
	}

	return answer
}

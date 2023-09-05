package puzzles

import (
	"strings"

	"github.com/anshulkanwar/aoc-go/internal"
)

func SolveDay2() {
	inputStr := internal.ReadInputFile("puzzles/day2/input.txt")

	var input [][]string

	for _, round := range strings.Split(inputStr, "\n") {
		input = append(input, strings.Split(round, " "))
	}

	selectPoint := make(map[string]int)
	selectPoint["X"] = 1
	selectPoint["Y"] = 2
	selectPoint["Z"] = 3

	score := 0

	for _, round := range input {
		opponent, you := round[0], round[1]

		score += selectPoint[you]

		if opponent == "A" {
			switch you {
			case "X":
				score += 3
			case "Y":
				score += 6
			case "Z":
				score += 0
			}
		} else if opponent == "B" {
			switch you {
			case "X":
				score += 0
			case "Y":
				score += 3
			case "Z":
				score += 6
			}
		} else if opponent == "C" {
			switch you {
			case "X":
				score += 6
			case "Y":
				score += 0
			case "Z":
				score += 3
			}
		}
	}

	println(score)
}

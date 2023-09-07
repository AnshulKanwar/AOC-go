package puzzles

import (
	"strings"

	"github.com/anshulkanwar/aoc-go/internal"
)

func SolveDay2() {
	var input [][]string

	inputStr := internal.ReadInputFile("puzzles/day2/input.txt")

	for _, round := range strings.Split(inputStr, "\n") {
		input = append(input, strings.Split(round, " "))
	}

	score := 0

	game := map[string]map[string]string{
		"A": {"Z": "paper", "Y": "rock", "X": "scissors"},
		"B": {"Z": "scissors", "Y": "paper", "X": "rock"},
		"C": {"Z": "rock", "Y": "scissors", "X": "paper"},
	}

	pointsOnChoice := map[string]int{"rock": 1, "paper": 2, "scissors": 3}
	pointsOnEnd := map[string]int{"Z": 6, "Y": 3, "X": 0}

	for _, round := range input {
		opponent, end := round[0], round[1]
		move := game[opponent][end]
		score += pointsOnChoice[move] + pointsOnEnd[end]
	}

	println(score)
}

package puzzles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anshulkanwar/aoc-go/internal"
)

func SolveDay1() {
	inputStr := internal.ReadInputFile("puzzles/day1/input.txt")

	var input [][]int

	for _, s := range strings.Split(inputStr, "\n\n") {
		caloriesStr := strings.Split(s, "\n")

		calories := make([]int, len(caloriesStr))

		for i, c := range caloriesStr {
			calorie, _ := strconv.Atoi(c)
			calories[i] = calorie
		}

		input = append(input, calories)
	}

	maxCalorie := 0

	for _, elf := range input {
		sum := 0
		for _, calorie := range elf {
			sum += calorie
		}
		maxCalorie = max(maxCalorie, sum)
	}

	fmt.Println(maxCalorie)
}

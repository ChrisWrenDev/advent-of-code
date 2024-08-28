package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ChrisWrenDev/advent-of-code/utils"
)

func solvePart1(input string) int {
	floor := 0

	for _, char := range input {
		if char == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return floor
}

func solvePart2(input string) int {
	floor := 0
	for position, char := range input {
		if floor == -1 {
			return position
		}
		if char == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return len(input)
}

func main() {
	part := utils.SolutionPartFlag()

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Failed to read input file:", err)
	}

	input := strings.TrimSpace(string(data))

	switch part {
	case "1":
		fmt.Printf("Part 1: %d\n", solvePart1(input))
	case "2":
		fmt.Printf("Part 2: %d\n", solvePart2(input))
	default:
		fmt.Printf("Part 1: %d\n", solvePart1(input))
		fmt.Printf("Part 2: %d\n", solvePart2(input))
	}
}

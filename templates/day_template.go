package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ChrisWrenDev/advent-of-code/utils"
)

func solvePart1(input string) int {
	// Your part 1 solution logic here
	return 0
}

func solvePart2(input string) int {
	// Your part 2 solution logic here
	return 0
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

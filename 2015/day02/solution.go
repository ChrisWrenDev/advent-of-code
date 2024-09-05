package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ChrisWrenDev/advent-of-code/utils"
)

func solvePart1(input []string) int64 {
	var paperSqFt int64
	for _, line := range input {
		output := strings.Split(line, "x")

		length, _ := strconv.ParseInt(output[0], 10, 64)
		width, _ := strconv.ParseInt(output[1], 10, 64)
		height, _ := strconv.ParseInt(output[2], 10, 64)

		var areas []int64
		areas = append(areas, 2*length*width)
		areas = append(areas, 2*width*height)
		areas = append(areas, 2*height*length)

		var surfaceArea int64
		minNum := areas[0]
		for _, num := range areas {
			if num < minNum {
				minNum = num
			}
			surfaceArea += num
		}

		paperSqFt = paperSqFt + surfaceArea + (minNum / 2)
	}
	return paperSqFt
}

func solvePart2(input []string) int64 {
	var ribbonFt int64
	for _, line := range input {
		output := strings.Split(line, "x")

		length, _ := strconv.ParseInt(output[0], 10, 64)
		width, _ := strconv.ParseInt(output[1], 10, 64)
		height, _ := strconv.ParseInt(output[2], 10, 64)

		ribbonBowFt := length * width * height

		var lengths []int64
		lengths = append(lengths, 2*(length+width))
		lengths = append(lengths, 2*(width+height))
		lengths = append(lengths, 2*(height+length))

		ribbonLength := lengths[0]
		for _, num := range lengths {
			if num < ribbonLength {
				ribbonLength = num
			}
		}

		ribbonFt = ribbonFt + ribbonLength + ribbonBowFt
	}
	return ribbonFt
}

func main() {
	part := utils.SolutionPartFlag()

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input file:", err)
	}

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

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

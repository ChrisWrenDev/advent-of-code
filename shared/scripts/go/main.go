package main

import (
	"fmt"

	"github.com/ChrisWrenDev/advent-of-code/scripts"
)

func main() {
	day, year, generateTemplate, getInput, getPrompt, checkCookie := scripts.ParseFlags()

	dayDir := fmt.Sprintf("%d/day%02d", year, day)

	if checkCookie {
		hasCookie := scripts.CheckAocCookie()
		if hasCookie {
			fmt.Println("Add AOC Session Cookie to .session file")
		} else {
			fmt.Println("AOC Session Cookie found in .session file")
		}
	}

	if getPrompt {
		body := scripts.GetAocPrompt(day, year)
		scripts.CreateAocFile(dayDir, "prompt.txt", body)
		fmt.Printf("Prompt for %d Day %d saved.\n", year, day)
	}

	if getInput {
		body := scripts.GetAocInput(day, year)
		scripts.CreateAocFile(dayDir, "input.txt", body)
		fmt.Printf("Input for %d Day %d saved.\n", year, day)
	}

	if generateTemplate {
		scripts.GenerateAocTemplate(dayDir)
		fmt.Println("Template created at: ", dayDir)
	}
}

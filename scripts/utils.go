package scripts

import (
	"flag"
	"log"
	"time"
)

func ParseFlags() (day, year int, generateTemplate, getInput, getPrompt bool) {
	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "Specify the day of the challenge")
	flag.IntVar(&year, "year", today.Year(), "Specify the year (default: 2024)")
	flag.BoolVar(&generateTemplate, "generate-template", false, "Generate template files for the day")
	flag.BoolVar(&getInput, "get-input", false, "Fetch input for the specified day")
	flag.BoolVar(&getPrompt, "get-prompt", false, "Fetch the prompt for the specified day")

	flag.Parse()

	if day > 25 || day < 1 {
		log.Fatalf("day out of range 1-25: %d", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	if year > today.Year() {
		log.Fatalf("year is after %d: %d", today.Year(), year)
	}

	return day, year, generateTemplate, getInput, getPrompt
}

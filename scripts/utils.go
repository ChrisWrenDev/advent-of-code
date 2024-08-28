package scripts

import (
	"flag"
	"log"
	"time"
)

func ParseFlags() (day, year int, generateTemplate, getInput, getPrompt, checkCookie bool) {
	today := time.Now()
	flag.IntVar(&day, "day", 0, "Specify the day of the challenge")
	flag.IntVar(&year, "year", 0, "Specify the year (default: 2024)")
	flag.BoolVar(&generateTemplate, "generate-template", false, "Generate template files for the day")
	flag.BoolVar(&getInput, "get-input", false, "Fetch input for the specified day")
	flag.BoolVar(&getPrompt, "get-prompt", false, "Fetch the prompt for the specified day")
	flag.BoolVar(&checkCookie, "check-cookie", false, "Check AOC session cookie has been set")

	flag.Parse()

	if day != 0 && (day > 25 || day < 1) {
		log.Fatalf("day out of range 1-25: %d", day)
	}

	if day != 0 && year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	if day != 0 && year > today.Year() {
		log.Fatalf("year is after %d: %d", today.Year(), year)
	}

	return day, year, generateTemplate, getInput, getPrompt, checkCookie
}

package utils

import (
	"flag"
	"fmt"
)

func SolutionPartFlag() string {
	part := flag.String("part", "0", "Specify which part to solve: 1 or 2")
	flag.Parse()

	if *part != "1" && *part != "2" {
		fmt.Println("Invalid part specified. Use -part=1 or -part=2.")
		return ""
	}

	return *part
}

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"aoc/internal/runner"
	_ "aoc/years" // Register all solutions
)

func main() {
	// Default to current year/day during December
	now := time.Now()
	defaultYear := now.Year()
	defaultDay := 0
	if now.Month() == time.December && now.Day() <= 25 {
		defaultDay = now.Day()
	}

	year := flag.Int("year", defaultYear, "Year to run (e.g., 2025)")
	day := flag.Int("day", defaultDay, "Day to run (1-25), 0 for all days")
	part := flag.Int("part", 0, "Part to run (1 or 2), 0 for both")
	inputFile := flag.String("input", "", "Custom input file (optional)")
	flag.Parse()

	if *day < 0 || *day > 25 {
		fmt.Fprintf(os.Stderr, "Error: day must be between 0 and 25\n")
		os.Exit(1)
	}

	cfg := runner.Config{
		Year:      *year,
		Day:       *day,
		Part:      *part,
		InputFile: *inputFile,
	}

	if err := runner.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

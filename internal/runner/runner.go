package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Solution represents a single day's solution
type Solution interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}

// Config holds the CLI configuration
type Config struct {
	Year      int
	Day       int
	Part      int
	InputFile string
}

// Registry holds all registered solutions
var registry = make(map[int]map[int]Solution)

// Register adds a solution to the registry
func Register(year, day int, sol Solution) {
	if registry[year] == nil {
		registry[year] = make(map[int]Solution)
	}
	registry[year][day] = sol
}

// Run executes solutions based on the config
func Run(cfg Config) error {
	yearSolutions, ok := registry[cfg.Year]
	if !ok {
		return fmt.Errorf("no solutions registered for year %d", cfg.Year)
	}

	if cfg.Day == 0 {
		// Run all days for the year
		for day := 1; day <= 25; day++ {
			if sol, exists := yearSolutions[day]; exists {
				if err := runDay(cfg.Year, day, sol, cfg.Part, cfg.InputFile); err != nil {
					fmt.Printf("Day %02d: Error - %v\n", day, err)
				}
			}
		}
		return nil
	}

	sol, ok := yearSolutions[cfg.Day]
	if !ok {
		return fmt.Errorf("no solution registered for year %d day %d", cfg.Year, cfg.Day)
	}

	return runDay(cfg.Year, cfg.Day, sol, cfg.Part, cfg.InputFile)
}

func runDay(year, day int, sol Solution, part int, inputFile string) error {
	input, err := loadInput(year, day, inputFile)
	if err != nil {
		return fmt.Errorf("failed to load input: %w", err)
	}

	fmt.Printf("=== Year %d Day %02d ===\n", year, day)

	if part == 0 || part == 1 {
		start := time.Now()
		result, err := sol.Part1(input)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("  Part 1: Error - %v\n", err)
		} else {
			fmt.Printf("  Part 1: %s (%v)\n", result, elapsed)
		}
	}

	if part == 0 || part == 2 {
		start := time.Now()
		result, err := sol.Part2(input)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("  Part 2: Error - %v\n", err)
		} else {
			fmt.Printf("  Part 2: %s (%v)\n", result, elapsed)
		}
	}

	return nil
}

func loadInput(year, day int, customPath string) (string, error) {
	var path string
	if customPath != "" {
		path = customPath
	} else {
		path = filepath.Join("years", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day), "input.txt")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

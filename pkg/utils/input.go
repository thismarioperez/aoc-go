package utils

import (
	"strconv"
	"strings"
)

// Lines splits input into lines, trimming whitespace and removing empty lines
func Lines(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}
	return strings.Split(input, "\n")
}

// LinesWithBlanks splits input into lines, preserving empty lines
func LinesWithBlanks(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

// Paragraphs splits input by blank lines
func Paragraphs(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}

// Ints parses all integers from a string (one per line or space-separated)
func Ints(input string) []int {
	var result []int
	for _, line := range Lines(input) {
		for _, field := range strings.Fields(line) {
			if n, err := strconv.Atoi(field); err == nil {
				result = append(result, n)
			}
		}
	}
	return result
}

// IntLines parses each line as a single integer
func IntLines(input string) []int {
	lines := Lines(input)
	result := make([]int, 0, len(lines))
	for _, line := range lines {
		if n, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
			result = append(result, n)
		}
	}
	return result
}

// Grid parses input into a 2D byte grid
func Grid(input string) [][]byte {
	lines := Lines(input)
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

// RuneGrid parses input into a 2D rune grid (for unicode support)
func RuneGrid(input string) [][]rune {
	lines := Lines(input)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// CSV parses comma-separated values from a single line
func CSV(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}

// CSVInts parses comma-separated integers from a single line
func CSVInts(input string) []int {
	parts := CSV(input)
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		if n, err := strconv.Atoi(strings.TrimSpace(p)); err == nil {
			result = append(result, n)
		}
	}
	return result
}

package day04

import (
	"fmt"
	"strconv"

	"aoc/internal/runner"
	"aoc/pkg/utils"
)

func init() {
	runner.Register(2025, 4, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	total := 0

	grid, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	if len(grid) == 0 {
		return "", fmt.Errorf("no grid found")
	}

	points := findPointsInGrid(grid)

	for _, point := range points {
		count := 0

		for _, neighbor := range point.Neighbors8() {
			x, ok := utils.GridAt(grid, neighbor)
			if ok && x == byte('@') {
				count++
			}
		}

		if count < 4 {
			total++
		}
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	grid, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	total := 0
	canContinue := true

	for canContinue {
		grid, total, canContinue = processGrid(grid, total)
	}

	return strconv.Itoa(total), nil
}

func parseInput(input string) ([][]byte, error) {
	grid := utils.Grid(input)

	return grid, nil
}

func findPointsInGrid(grid [][]byte) []utils.Point {
	return utils.FindInGrid(grid, byte('@'))
}

func processGrid(grid [][]byte, total int) ([][]byte, int, bool) {
	points := findPointsInGrid(grid)
	var toRemove []utils.Point

	for _, point := range points {
		count := 0

		for _, neighbor := range point.Neighbors8() {
			x, ok := utils.GridAt(grid, neighbor)
			if ok && x == byte('@') {
				count++
			}
		}

		if count < 4 {
			toRemove = append(toRemove, point)
		}
	}

	for _, point := range toRemove {
		grid[point.Y][point.X] = byte('.')
	}

	total += len(toRemove)
	canContinue := len(toRemove) > 0

	return grid, total, canContinue
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

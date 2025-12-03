package day04

import (
	"fmt"
	"strconv"

	"aoc/internal/runner"
	"aoc/pkg/utils"
)

func init() {
	runner.Register(2024, 4, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	SEARCH_STR := "XMAS"
	total := 0

	grid, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	width, height := utils.GridBounds(grid)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for _, dir := range utils.AllDirs {
				dx, dy := dir.X, dir.Y

			}
		}
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	return strconv.Itoa(9), nil
}

func parseInput(input string) ([][]byte, error) {
	grid := utils.Grid(input)

	return grid, nil
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

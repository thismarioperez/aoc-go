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

	// Search for all "X" characters
	startPoints := utils.FindInGrid(grid, []byte(SEARCH_STR)[0])

	// fmt.Printf("Start points: %v\n", startPoints)

	for _, startPoint := range startPoints {
	dirLoop:
		for _, dir := range utils.AllDirs {
			for i := 0; i < len(SEARCH_STR); i++ {
				letter, ok := utils.GridAt(grid, startPoint.Add(dir.Scale(i)))
				if !ok || letter != SEARCH_STR[i] {
					continue dirLoop
				}
			}
			total++
			// fmt.Printf("Found valid sequence at point: %v in direction: %v\n", startPoint, dir)
		}
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	SEARCH_STR := "MAS"
	total := 0

	grid, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	// Search for all "A" characters
	startPoints := utils.FindInGrid(grid, []byte(SEARCH_STR)[1])
	// fmt.Printf("Start points: %v\n", startPoints)

	for _, center := range startPoints {
		ul, ok1 := utils.GridAt(grid, center.Add(utils.UpLeft))
		ur, ok2 := utils.GridAt(grid, center.Add(utils.UpRight))
		dl, ok3 := utils.GridAt(grid, center.Add(utils.DownLeft))
		dr, ok4 := utils.GridAt(grid, center.Add(utils.DownRight))

		if !ok1 || !ok2 || !ok3 || !ok4 {
			// Out of bounds
			continue
		}

		isDiag1Valid := ul == SEARCH_STR[0] && dr == SEARCH_STR[2] || ul == SEARCH_STR[2] && dr == SEARCH_STR[0]
		isDiag2Valid := ur == SEARCH_STR[0] && dl == SEARCH_STR[2] || ur == SEARCH_STR[2] && dl == SEARCH_STR[0]

		if isDiag1Valid && isDiag2Valid {
			total++
			// fmt.Printf("Found valid sequence at point: %v\n", center)
		}
	}

	return strconv.Itoa(total), nil
}

func parseInput(input string) ([][]byte, error) {
	grid := utils.Grid(input)

	return grid, nil
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

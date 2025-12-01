package day01

import (
	"fmt"
	"strconv"

	"aoc/internal/runner"
	"aoc/pkg/utils"
)

func init() {
	runner.Register(2025, 1, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	position := 50
	hits := 0

	deltas, err := parseInput(input)

	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, delta := range deltas {
		position = position + delta

		if position%100 == 0 {
			hits++
		}

		if position < 0 {
			position += 100
		} else if position > 99 {
			position -= 100
		}
	}

	return strconv.Itoa(hits), nil
}

func (s *Solution) Part2(input string) (string, error) {
	position := 50
	hits := 0

	deltas, err := parseInput(input)

	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, delta := range deltas {
		for i := 0; i < utils.Abs(delta); i++ {
			position = position + utils.Sign(delta)
			if position%100 == 0 {
				hits++
			}
			if position < 0 {
				position += 100
			} else if position > 99 {
				position -= 100
			}
		}
	}

	return strconv.Itoa(hits), nil
}

type Direction string

const (
	DirectionLeft  Direction = "L"
	DirectionRight Direction = "R"
)

func parseInput(input string) ([]int, error) {
	lines := utils.Lines(input)
	deltas := make([]int, 0, len(lines))

	for _, line := range lines {
		direction := string(line[0])
		delta, err := strconv.Atoi(line[1:])
		if direction == string(DirectionLeft) {
			delta = -delta
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse delta: %w", err)
		}
		deltas = append(deltas, delta)
	}

	return deltas, nil
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

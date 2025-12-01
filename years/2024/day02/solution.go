package day02

import (
	"aoc/internal/runner"
	"aoc/pkg/utils"
	"strconv"
)

func init() {
	runner.Register(2024, 2, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	lines := parseInput(input)

	var safeCount int

	for _, line := range lines {
		fields := utils.Ints(line)
		if detectIsSafe(fields) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount), nil
}

func (s *Solution) Part2(input string) (string, error) {

	lines := parseInput(input)

	var safeCount int

	for _, line := range lines {
		fields := utils.Ints(line)
		if detectIsSafeWithDamping(fields) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount), nil
}

func parseInput(input string) []string {
	lines := utils.Lines(input)

	return lines
}

func detectIsSafe(fields []int) bool {
	isSafe := true
	var isIncreasing *bool
	currentField := 0
	nextField := 0
	difference := 0

	for idx := 0; idx < len(fields)-1; idx++ {
		if !isSafe {
			break
		}
		currentField = fields[idx]
		nextField = fields[idx+1]
		difference = utils.Abs(nextField - currentField)

		// Set increasing flag
		if isIncreasing == nil {
			_isIncreasing := nextField > currentField
			isIncreasing = &_isIncreasing
		} else if *isIncreasing && nextField < currentField {
			isSafe = false
		} else if !*isIncreasing && nextField > currentField {
			isSafe = false
		}

		// Check if difference is in valid range (1-3)
		if difference < 1 || difference > 3 {
			isSafe = false
		}
	}

	return isSafe
}

func detectIsSafeWithDamping(fields []int) bool {
	isSafe := detectIsSafe(fields)
	var dampedFields []int

	// Short circuit if the fields are safe
	if isSafe {
		return isSafe
	}

	if !isSafe {
		for idx := 0; idx < len(fields); idx++ {
			// Short circuit if a previous iteration found a safe sequence
			if isSafe {
				break
			}
			// Copy fields then omit current field and check if the sequence is safe
			dampedFields = make([]int, 0, len(fields)-1)
			dampedFields = append(dampedFields, fields[:idx]...)
			dampedFields = append(dampedFields, fields[idx+1:]...)

			isSafe = detectIsSafe(dampedFields)
		}
	}

	return isSafe
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

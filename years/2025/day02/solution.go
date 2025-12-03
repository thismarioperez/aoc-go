package day02

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/internal/runner"
	"aoc/pkg/utils"
)

func init() {
	runner.Register(2025, 2, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	total := 0

	idRanges, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			if isRepeatingTwice(strconv.Itoa(id)) {
				// fmt.Printf("%s\n", fmt.Sprintf("found repeating id: %d", id))
				total += id
			}
		}

	}

	return strconv.Itoa(total), nil

}

func (s *Solution) Part2(input string) (string, error) {
	total := 0

	idRanges, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			if isRepeating(strconv.Itoa(id)) {
				// fmt.Printf("%s\n", fmt.Sprintf("found repeating id: %d", id))
				total += id
			}
		}

	}

	return strconv.Itoa(total), nil
}

type IDRange struct {
	Start int
	End   int
}

func parseInput(input string) ([]IDRange, error) {
	values := utils.CSV(input)
	idRanges := make([]IDRange, 0, len(values))

	for _, value := range values {
		parts := strings.Split(value, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse start: %w", err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse end: %w", err)
		}

		idRanges = append(idRanges, IDRange{start, end})
	}

	return idRanges, nil
}

func isRepeatingTwice(s string) bool {
	// Sequence is repeating twice if first half equals second half
	if len(s)%2 == 0 {
		half := len(s) / 2

		first := s[:half]
		second := s[half:]

		return first == second
	}

	return false
}

func isRepeating(s string) bool {
	// double the string, drop the first and last characters,
	// then check if the original string is contained in the doubled string
	// this proves that the pattern repeats since the original string is contained in the doubled string
	doubled := s + s
	dropped := doubled[1 : len(doubled)-1]

	return strings.Contains(dropped, s)
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

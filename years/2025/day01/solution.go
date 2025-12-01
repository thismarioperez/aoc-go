package day01

import (
	"sort"
	"strconv"
	"strings"

	"aoc/internal/runner"
	"aoc/pkg/utils"
)

func init() {
	runner.Register(2025, 1, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	left, right := parseInput(input)

	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i := range left {
		total += utils.Abs(left[i] - right[i])
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	left, right := parseInput(input)

	// Count occurrences in right list
	counts := make(map[int]int)
	for _, n := range right {
		counts[n]++
	}

	// Calculate similarity score
	total := 0
	for _, n := range left {
		total += n * counts[n]
	}

	return strconv.Itoa(total), nil
}

func parseInput(input string) ([]int, []int) {
	lines := utils.Lines(input)
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			l, _ := strconv.Atoi(fields[0])
			r, _ := strconv.Atoi(fields[1])
			left = append(left, l)
			right = append(right, r)
		}
	}

	return left, right
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

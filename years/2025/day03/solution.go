package day03

import (
	"aoc/internal/runner"
	"aoc/pkg/utils"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	runner.Register(2025, 3, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	total := 0
	banks, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, bank := range banks {
		result := maxNDigitNumberString(bank, 2)
		// fmt.Printf("result: %s\n", result)
		num, err := strconv.Atoi(result)
		if err != nil {
			return "", fmt.Errorf("failed to parse number: %w from string %s", err, result)
		}

		total += num
	}

	// for _, bank := range banks {
	// 	result := []int{}

	// 	// Generate all possible two digit numbers from the bank
	// 	for i := 0; i < len(bank)-1; i++ {
	// 		for j := i + 1; j < len(bank); j++ {
	// 			s := bank[i] + bank[j]
	// 			n, err := strconv.Atoi(s)
	// 			if err != nil {
	// 				return "", fmt.Errorf("failed to parse number: %w from string %s", err, s)
	// 			}
	// 			result = append(result, n)
	// 		}
	// 	}

	// 	max := utils.Max(result...)
	// 	total += max
	// }

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	total := 0
	banks, err := parseInput(input)

	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, bank := range banks {
		result := maxNDigitNumberString(bank, 12)
		// fmt.Printf("result: %s\n", result)
		num, err := strconv.Atoi(result)
		if err != nil {
			return "", fmt.Errorf("failed to parse number: %w from string %s", err, result)
		}

		total += num
	}

	return strconv.Itoa(total), nil
}

func parseInput(input string) ([][]int, error) {
	lines := utils.Lines(input)
	banks := [][]int{}

	for _, line := range lines {
		digits := strings.Join(strings.Split(line, ""), " ")
		bank := utils.Ints(digits)
		banks = append(banks, bank)
	}

	return banks, nil
}

func maxNDigitNumberString(bank []int, n int) string {
	result := ""
	startIdx := 0 // Index of the first digit in bank to consider

	// Loop over the n range (eg 2 for 2 digit number, 12 for 12 digit number, etc.)
	for currentDigitIdx := 0; currentDigitIdx < n; currentDigitIdx++ {
		// Calculate endIdx for each iteration
		// Leave enough room for remaining digits
		// First iteration should equal the length of the bank - 1
		// Second iteration should equal length of bank - 2, etc...
		endIdx := len(bank) - (n - currentDigitIdx - 1)
		maxIdx := startIdx
		for idx := startIdx; idx < endIdx; idx++ {
			curVal := bank[idx]
			maxVal := bank[maxIdx]
			if curVal > maxVal {
				maxIdx = idx
			}
		}
		current := strconv.Itoa(bank[maxIdx])
		result += current
		startIdx = maxIdx + 1
	}

	return result

}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

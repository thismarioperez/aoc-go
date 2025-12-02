package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc/internal/runner"
)

func init() {
	runner.Register(2024, 3, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
	commands, err := parseInput1(input)
	total := 0
	inputs := []int{0, 0}

	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, command := range commands {
		matches := regexp.MustCompile(INPUTS_EXPRESSION).FindAllString(command, -1)
		for i, match := range matches {
			value, err := strconv.Atoi(match)
			if err != nil {
				return "", fmt.Errorf("failed to parse input: %w", err)
			}
			inputs[i] = value
		}

		total += inputs[0] * inputs[1]
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	commands, err := parseInput2(input)
	total := 0
	inputs := []int{0, 0}
	enabled := true

	if err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	for _, command := range commands {
		if regexp.MustCompile(DONT_EXPRESSION).MatchString(command) {
			enabled = false
		} else if regexp.MustCompile(DO_EXPRESSION).MatchString(command) {
			enabled = true
		} else if regexp.MustCompile(MUL_EXPRESSION).MatchString(command) && enabled {
			matches := regexp.MustCompile(INPUTS_EXPRESSION).FindAllString(command, -1)
			for i, match := range matches {
				value, err := strconv.Atoi(match)

				if err != nil {
					return "", fmt.Errorf("failed to parse input: %w", err)
				}

				inputs[i] = value
			}

			total += inputs[0] * inputs[1]
		}
	}

	return strconv.Itoa(total), nil
}

var MUL_EXPRESSION = `mul\(\d{1,3},\d{1,3}\)`
var INPUTS_EXPRESSION = `\d{1,3}`
var DO_EXPRESSION = `do\(\)`
var DONT_EXPRESSION = `don'?t\(\)`
var patterns = []string{
	MUL_EXPRESSION,
	DO_EXPRESSION,
	DONT_EXPRESSION,
}
var COMBINED_EXPRESSION = strings.Join(patterns, "|")

func parseInput1(input string) ([]string, error) {
	commands := regexp.MustCompile(MUL_EXPRESSION).FindAllString(input, -1)

	return commands, nil
}

func parseInput2(input string) ([]string, error) {
	commands := regexp.MustCompile(COMBINED_EXPRESSION).FindAllString(input, -1)

	return commands, nil
}

// Ensure Solution implements runner.Solution
var _ runner.Solution = (*Solution)(nil)

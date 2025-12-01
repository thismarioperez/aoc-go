package day01

import "testing"

const exampleInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {
	s := &Solution{}
	result, err := s.Part1(exampleInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "3"
	if result != expected {
		t.Errorf("Part1() = %q, want %q", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := &Solution{}
	result, err := s.Part2(exampleInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "6"
	if result != expected {
		t.Errorf("Part2() = %q, want %q", result, expected)
	}
}

// Benchmarks
func BenchmarkPart1(b *testing.B) {
	s := &Solution{}
	for i := 0; i < b.N; i++ {
		s.Part1(exampleInput)
	}
}

func BenchmarkPart2(b *testing.B) {
	s := &Solution{}
	for i := 0; i < b.N; i++ {
		s.Part2(exampleInput)
	}
}

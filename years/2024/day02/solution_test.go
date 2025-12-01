package day02

import "testing"

const exampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	s := &Solution{}
	result, err := s.Part1(exampleInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "2"
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

	expected := "4"
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

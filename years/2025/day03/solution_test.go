package day03

import "testing"

const exampleInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	s := &Solution{}
	result, err := s.Part1(exampleInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "357"
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

	expected := "3121910778619"
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

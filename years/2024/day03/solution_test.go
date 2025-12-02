package day03

import "testing"

const exampleInput1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const exampleInput2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart1(t *testing.T) {
	s := &Solution{}
	result, err := s.Part1(exampleInput1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "161"
	if result != expected {
		t.Errorf("Part1() = %q, want %q", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := &Solution{}
	result, err := s.Part2(exampleInput2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "48"
	if result != expected {
		t.Errorf("Part2() = %q, want %q", result, expected)
	}
}

// Benchmarks
func BenchmarkPart1(b *testing.B) {
	s := &Solution{}
	for i := 0; i < b.N; i++ {
		s.Part1(exampleInput1)
	}
}

func BenchmarkPart2(b *testing.B) {
	s := &Solution{}
	for i := 0; i < b.N; i++ {
		s.Part2(exampleInput2)
	}
}

package day02

import "testing"

const exampleInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	s := &Solution{}
	result, err := s.Part1(exampleInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "1227775554"
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

	expected := "4174379265"
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

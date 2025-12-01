package utils

import "cmp"

// Abs returns the absolute value of x
func Abs[T ~int | ~int64 | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of the given values
func Min[T cmp.Ordered](vals ...T) T {
	if len(vals) == 0 {
		var zero T
		return zero
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the maximum of the given values
func Max[T cmp.Ordered](vals ...T) T {
	if len(vals) == 0 {
		var zero T
		return zero
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

// Sum returns the sum of all values
func Sum[T ~int | ~int64 | ~float64](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}

// GCD returns the greatest common divisor of a and b
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs(a)
}

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / GCD(a, b)
}

// LCMSlice returns the LCM of all values in the slice
func LCMSlice(vals []int) int {
	if len(vals) == 0 {
		return 0
	}
	result := vals[0]
	for _, v := range vals[1:] {
		result = LCM(result, v)
	}
	return result
}

// Sign returns -1, 0, or 1 based on the sign of x
func Sign[T ~int | ~int64 | ~float64](x T) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

// Clamp constrains x to be between lo and hi
func Clamp[T cmp.Ordered](x, lo, hi T) T {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}

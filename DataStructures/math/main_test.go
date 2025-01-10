package main

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{3, 5, 1},
		{23, 5, 1},
		{15, 25, 5},
		{56, 48, 8},
		{0, 7, 7},
		{18, 0, 18},
		{0, 0, 0}, // You might want to define the behavior for this case
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("GCD(%d, %d)", tc.a, tc.b)
		t.Run(name, func(t *testing.T) {
			result := gcd(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected GCD(%d, %d) to be %d, but got %d", tc.a, tc.b, tc.expected, result)
			}
		})
	}
}

package main

import "testing"

func Test_bst_validateBst(t *testing.T) {
	tests := []struct {
		input       []int
		expectedOut bool
	}{
		{[]int{6, 3, 9, 2, 4, 8, 10}, true},
		{[]int{6, 3, 9, 2, 4, 8, 10, 1}, true},
	}
	for _, s := range tests {
		b := new(bst)
		for _, x := range s.input {
			b.insert(x)
		}
		out := b.ValidateBst()
		if s.expectedOut != out {
			t.Errorf("test failed because got out %v instead of %v", out, s.expectedOut)
		}
	}
}

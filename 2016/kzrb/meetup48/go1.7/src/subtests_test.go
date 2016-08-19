package src

import (
	"fmt"
	"testing"
)

// START OMIT
func TestSubtests(t *testing.T) {
	tests := []struct {
		A, B int
		Sum  int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, tc := range tests {
		d := fmt.Sprint(tc.A, "+", tc.B)
		t.Run(d, func(t *testing.T) {
			r := tc.A + tc.B
			if r != tc.Sum {
				t.Errorf("Expected %d, actual %d", tc.Sum, r)
			}
		})
	}
}

// END OMIT

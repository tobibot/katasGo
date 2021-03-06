package kata_test

import (
	"testing"

	kata "github.com/tobibot/katasGo/sudokuSolutionValidator"
)

// kata "github.com/tobibot/katasGo/sudokuSolutionValidator"

var testTrue = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

var testFalse = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 0, 3, 4, 8},
	{1, 0, 0, 3, 4, 2, 5, 6, 0},
	{8, 5, 9, 7, 6, 1, 0, 2, 0},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 0, 1, 5, 3, 7, 2, 1, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 0, 0, 4, 8, 1, 1, 7, 9},
}

var testShouldBeFalse = [][]int{
	{5, 3, 4, 6, 7, 8, 1, 9, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

func TestValidateSolution(t *testing.T) {
	tests := []struct {
		name string
		m    [][]int
		want bool
	}{
		{"valid", testTrue, true},
		{"invalid", testFalse, false},
		{"shouldBeFalse", testShouldBeFalse, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kata.ValidateSolution(tt.m); got != tt.want {

				t.Errorf("ValidateSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

package kata_test

import (
	"reflect"
	"testing"

	kata "github.com/tobibot/katasGo/snail"
)

func TestSnail(t *testing.T) {
	tests := []struct {
		name     string
		snaipMap [][]int
		want     []int
	}{
		{"realy empty matrix", [][]int{{}}, []int{}},
		{"1x1", [][]int{{1}}, []int{1}},
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"4x4", [][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}, []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kata.Snail(tt.snaipMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Snail() = %v, want %v", got, tt.want)
			}
		})
	}
}

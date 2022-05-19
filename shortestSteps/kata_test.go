package kata_test

import (
	"testing"

	kata "github.com/tobibot/katasGo/shortestSteps"
)

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

func TestShortestStepsToNum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"1", 1, 0},
		{"12", 12, 4},
		{"16", 16, 4},
		{"71", 71, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kata.ShortestStepsToNum(tt.num); got != tt.want {
				t.Errorf("ShortestStepsToNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

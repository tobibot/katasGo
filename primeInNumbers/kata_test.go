package kata_test

import (
	"testing"

	kata "github.com/tobibot/katasGo/primeInNumbers"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"6", 6, "(2)(3)"},
		{"10", 10, "(2)(5)"},
		{"a", 79340, "(2**2)(5)(3967)"},
		{"a", 7793, "(7793)"},
		{"b", 7775460, "(2**2)(3**3)(5)(7)(11**2)(17)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kata.PrimeFactors(tt.n); got != tt.want {
				t.Errorf("PrimeFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}

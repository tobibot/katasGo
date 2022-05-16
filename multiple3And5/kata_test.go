package kata_test

import (
	"testing"

	"github.com/tobibot/katasGo/multiple3And5"
	// 	. "github.com/onsi/ginkgo"
	// 	. "github.com/onsi/gomega"
)

// var _ = Describe("Multiples of 3 and 5", func() {

// 	It("should handle basic cases", func() {
// 		Expect(Multiple3And5(10)).To(Equal(23))
// 	})
// })

func TestMultiple3And5(t *testing.T) {
	
	tests := []struct {
		name string
		number int
		want int
	}{
		{"10", 10, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kata.Multiple3And5(tt.number); got != tt.want {
				t.Errorf("Multiple3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}

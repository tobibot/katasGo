package isbn10validation

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ISBN Validator", func() {
	It("should recognize valid ISBNs", func() {
		Expect(ValidISBN10("1112223339")).To(Equal(true), "1112223339 is valid")
		Expect(ValidISBN10("048665088X")).To(Equal(true), "048665088X is valid")
		Expect(ValidISBN10("1293000000")).To(Equal(true), "1293000000 is valid")
		Expect(ValidISBN10("1234554321")).To(Equal(true), "1234554321 is valid")
	})

	It("should recognize invalid ISBNs", func() {
		Expect(ValidISBN10("1234512345")).To(Equal(false), "1234512345 is not valid")
		Expect(ValidISBN10("1293")).To(Equal(false), "1293 is not valid")
		Expect(ValidISBN10("X123456788")).To(Equal(false), "X123456788 is not valid")
		Expect(ValidISBN10("ABCDEFGHIJ")).To(Equal(false), "ABCDEFGHIJ is not valid")
		Expect(ValidISBN10("XXXXXXXXXX")).To(Equal(false), "XXXXXXXXXX is not valid")
	})
})

func TestValidISBN10(t *testing.T) {
	tests := []struct {
		name string
		isbn string
		want bool
	}{
		// {"1112223339 is valid", "1112223339", true},
		// {"048665088X is valid", "048665088X", true},
		// {"1293000000 is valid", "1293000000", true},
		// {"1234554321 is valid", "1234554321", true},
		{"X123456788 is valid", "048665088x", true},

		// {"1234512345 is not valid", "1234512345", false},
		// {"1293 is not valid", "1293", false},
		// {"X123456788 is not valid", "X123456788", false},

		// {"ABCDEFGHIJ is not valid", "ABCDEFGHIJ", false},
		// {"XXXXXXXXXX is not valid", "XXXXXXXXXX", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidISBN10(tt.isbn); got != tt.want {
				t.Errorf("ValidISBN10() = %v, want %v", got, tt.want)
			}
		})
	}
}

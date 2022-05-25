package isbn10validation

import (
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

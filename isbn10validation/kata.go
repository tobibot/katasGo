package isbn10validation

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	digitRegexp = regexp.MustCompile(`\d`)
)

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	s := strings.Split(isbn, "")
	sum := 0
	for i, r := range s {		
		if digitRegexp.MatchString(r) {
			toAdd, _ := strconv.Atoi(r)
			sum += toAdd * (i + 1)			
		} else if i== 9 && (r == "X" || r == "x") {
			sum += (i + 1) * 10
		} else {
			return false
		}
	}
	return sum%11 == 0
}

package kata

import (
	"fmt"
	"math"
	"sort"
)

var (
	primeFactors map[int]int
	goal         int
)

func PrimeFactors(n int) string {
	primeFactors = make(map[int]int)
	goal = n
	getNextPrimeFactor(n)
	return prettify()

}
func prettify() string {
	result := ""

	// Todo : sort
	keys := make([]int, 0, len(primeFactors))

	for k := range primeFactors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		v := primeFactors[k]
		// for k, v := range primeFactors {
		if v == 1 {
			result += fmt.Sprintf("(%v)", k)
		} else {
			result += fmt.Sprintf("(%v**%v)", k, v)
		}
	}
	return result
}

func resultSoFar() int {
	product := 0
	for k, v := range primeFactors {
		product += k * v
	}
	return product
}

func getNextPrimeFactor(number int) {
	if number == 1 {
		return
	}
	lmt := int(math.Ceil(math.Sqrt((float64(number))))) + 1
	for i := 2; i <= lmt; i++ {
		if r := resultSoFar(); r == goal {
			return
		}
		if number%i == 0 {
			add(i)
			getNextPrimeFactor(number / i)
			return
		}
	}
	add(number)
}

func add(i int) {
	if v, found := primeFactors[i]; found {
		primeFactors[i] = v + 1
	} else {
		primeFactors[i] = 1
	}
}

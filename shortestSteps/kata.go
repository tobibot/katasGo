package kata

func ShortestStepsToNum(n int) int {
	// your code here

	if n < 1 || n > 1000000 {
		return -1
	}
	if n%2 == 0 {
		return ShortestStepsToNum(n/2) + 1
	} else if n > 1 {
		return ShortestStepsToNum(n-1) + 1
	} else {
		return 0
	}

}

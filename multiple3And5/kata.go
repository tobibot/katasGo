package kata

func Multiple3And5(number int) int {

	if number < 0 {
		return 0
	}
	sum := 0
	for i := 0; i < number; i++ {	
		if i%3 == 0 {
			sum += i
			continue
		}
		if i%5 == 0 {
			sum += i
			continue
		}
		// if-statements can be OR'ed
	}
	return sum

}

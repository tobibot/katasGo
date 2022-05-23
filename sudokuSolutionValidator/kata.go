package kata

import "sort"

// kata "github.com/tobibot/katasGo/sudokuSolutionValidator"

func ValidateSolution(m [][]int) bool {
	var row1stQuadrant, row2ndQuadrant, row3rdQuadrant [9]int

	for i, row := range m {

		var actRowArray [9]int
		copy(actRowArray[:], row[0:9]) // syntax sieht falsch aus, ist aber so
		rowsToCheck := [][9]int{actRowArray}

		for j := (i % 3) * 3; j < ((i%3)+1)*3; j++ { // Das war die Hölle
			row1stQuadrant[j] = row[j%3]
			row2ndQuadrant[j] = row[(j%3)+3]
			row3rdQuadrant[j] = row[(j%3)+6]
		}

		if (i+1)%3 == 0 {
			rowsToCheck = append(rowsToCheck, row1stQuadrant, row2ndQuadrant, row3rdQuadrant)
		}
		if r := check(rowsToCheck); !r {
			return false
		}
	}

	return true
}

func check(m [][9]int) bool {
	for _, row := range m {
		if o := rowOkay(row); !o {
			return o
		}
	}
	return true
}

func rowOkay(r [9]int) bool {
	sort.Ints(r[:])
	if r[0] == 0 {
		return false
	}

	for i := 0; i < 7; i++ {
		if r[i]+1 != r[i+1] {
			return false
		}
	}
	return true
}

package kata

const (
	VISITED = 2147483647
)

var (
	yDir         int
	xDir         int
	curX         int
	curY         int
	limit        int
	resultLength int
	snailMap     [][]int
)

func initVars() {
	yDir = 0
	xDir = 1
	curX = 0
	curY = 0
}

func Snail(snaipMap [][]int) []int {
	
	if len(snaipMap) == 1 && len(snaipMap[0]) == 0 {
		return []int{}
	} else if len(snaipMap) == 1 {
		return []int{snaipMap[0][0]}
	}

	initVars()

	limit = len(snaipMap) - 1
	resultLength = len(snaipMap) * len(snaipMap)
	result := make([]int, 0)
	snailMap = snaipMap

	result = append(result, snaipMap[curY][curX])
	snailMap[curY][curX] = VISITED

	for {
		curX += xDir
		curY += yDir

		result = append(result, snaipMap[curY][curX])
		snaipMap[curY][curX] = VISITED
		if len(result) == resultLength {
			return result
		}

		if borderReached() || nextAlreadyVisited() {
			changeDirection()
		}
	}

}

func nextAlreadyVisited() bool {
	return snailMap[curY+yDir][curX+xDir] == VISITED
}

func borderReached() bool {
	return (curX+xDir > limit || curX+xDir < 0 || curY+yDir > limit || curY+yDir < 0)
}

func changeDirection() {
	if xDir == 1 && yDir == 0 {
		xDir = 0
		yDir = 1
	} else if xDir == 0 && yDir == 1 {
		xDir = -1
		yDir = 0
	} else if xDir == -1 && yDir == 0 {
		xDir = 0
		yDir = -1
	} else if xDir == 0 && yDir == -1 {
		xDir = 1
		yDir = 0
	}
}

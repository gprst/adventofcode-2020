package main

func getEncounteredTreesCount(rows []string) (treesCount uint16) {
	curPos := position{0, 0}
	for curPos.y != uint16(len(rows)) {
		if curPos.x >= patternLength {
			curPos.x -= patternLength
		}
		if rows[curPos.y][curPos.x] == byte('#') {
			treesCount++
		}
		curPos = calculatePosition(curPos, slope{3, 1})
	}
	return treesCount
}

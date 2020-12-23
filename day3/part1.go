package main

func getEncounteredTreesCount(rows []string) (treesCount uint16) {
	curPos := Position{0, 0}
	patternLength := uint16(len(rows[0]))
	for curPos.y != uint16(len(rows)) {
		if curPos.x >= patternLength {
			curPos.x -= patternLength
		}
		if rows[curPos.y][curPos.x] == byte('#') {
			treesCount++
		}
		curPos.x += 3
		curPos.y++
	}
	return treesCount
}

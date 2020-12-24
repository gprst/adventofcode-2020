package main

func getTreesCountForSlope(slope slope, rows []string) uint16 {
	var treesCount uint16
	curPos := position{0, 0}
	lastRow := uint16(len(rows))
	for curPos.y < lastRow {
		if rows[curPos.y][curPos.x] == byte('#') {
			treesCount++
		}
		curPos = calculatePosition(curPos, slope)
	}
	return treesCount
}

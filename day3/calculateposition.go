package main

func calculatePosition(curPos position, slope slope) position {
	nextX := curPos.x + uint16(slope.right)
	nextY := curPos.y + uint16(slope.down)
	if nextX >= patternLength {
		nextX -= patternLength
	}
	return position{nextX, nextY}
}

package main

import "math"

func getPosition(sequence string, min int, max int, lowerHalf rune, upperHalf rune) int {
	for _, letter := range sequence {
		difference := int(math.Ceil(float64(max-min) / 2))
		if letter == lowerHalf {
			max -= difference
		} else if letter == upperHalf {
			min += difference
		}
	}
	return min
}

func getSeatID(boardingPass string) int {
	rowSequence := boardingPass[:7]
	columnSequence := boardingPass[7:]
	return getPosition(rowSequence, 0, 127, 'F', 'B')*8 + getPosition(columnSequence, 0, 7, 'L', 'R')
}

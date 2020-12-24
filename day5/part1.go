package main

import "math"

func getPosition(sequence string, min uint, max uint, lowerHalf rune, upperHalf rune) uint {
	for _, letter := range sequence {
		difference := uint(math.Ceil(float64(max-min) / 2))
		if letter == lowerHalf {
			max -= difference
		} else if letter == upperHalf {
			min += difference
		}
	}
	return min
}

func getSeatID(boardingPass string) uint {
	rowSequence := boardingPass[:7]
	columnSequence := boardingPass[7:]
	return getPosition(rowSequence, 0, 127, 'F', 'B')*8 + getPosition(columnSequence, 0, 7, 'L', 'R')
}

func getHighestSeatID(boardingPasses []string) (highestSeatID uint) {
	for _, boardingPass := range boardingPasses {
		seatID := getSeatID(boardingPass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

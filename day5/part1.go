package main

func getHighestSeatID(boardingPasses []string) (highestSeatID uint) {
	for _, boardingPass := range boardingPasses {
		seatID := getSeatID(boardingPass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID
}

package main

import "sort"

func getAllSortedSeatIDs(boardingPasses []string) (seatIDs []int) {
	for _, boardingPass := range boardingPasses {
		seatIDs = append(seatIDs, int(getSeatID(boardingPass)))
	}
	sort.Ints(seatIDs)
	return seatIDs
}

func getYourSeatID(boardingPasses []string) int {
	seatIDs := getAllSortedSeatIDs(boardingPasses)
	for i := 1; i < len(seatIDs); i++ {
		current := seatIDs[i]
		previous := seatIDs[i-1]
		if current == previous+2 {
			return current - 1
		}
	}
	return 0
}

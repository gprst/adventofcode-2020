package main

func getMultipliedEncounteredTreesCount(rows []string) uint {
	var result uint = 1
	slopes := [...]slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		result *= uint(getTreesCountForSlope(slope, rows))
	}
	return result
}

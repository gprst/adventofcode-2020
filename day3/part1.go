package main

func getEncounteredTreesCount(rows []string) (treesCount uint16) {
	return getTreesCountForSlope(slope{3, 1}, rows)
}

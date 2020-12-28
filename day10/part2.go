package main

func getAllAdapterTreeLeaves(nextNodesJoltage []int, tree [][]int, cache []int) (result int) {
	if len(nextNodesJoltage) == 0 {
		return 1
	}
	for _, nextJoltage := range nextNodesJoltage {
		if cache[nextJoltage] != 0 {
			result += cache[nextJoltage]
		} else {
			computedValue := getAllAdapterTreeLeaves(tree[nextJoltage], tree, cache)
			cache[nextJoltage] = computedValue
			result += computedValue
		}
	}
	return result
}

func mapAdapterTree(adapters []int, adaptersLength int) [][]int {
	contains := func(s []int, number int) bool {
		for _, element := range s {
			if element == number {
				return true
			}
		}
		return false
	}
	// Each index corresponds to an adapter value. Empty array means no adapter
	// with this value, and an array with numbers represents the adapter
	// children.
	maxJoltage := adapters[len(adapters)-1] + 3
	tree := make([][]int, maxJoltage)
	for joltage := 0; joltage <= maxJoltage; joltage++ {
		if contains(adapters, joltage) {
			possibleOuputs := []int{}
			for i := 1; i <= 3; i++ {
				output := joltage + i
				if contains(adapters, output) {
					possibleOuputs = append(possibleOuputs, output)
				}
			}
			tree[joltage] = possibleOuputs
		}

	}
	return tree
}

func getAdaptersCombinations(adapters []int) (combinations int) {
	adapters = append([]int{0}, adapters...)
	tree := mapAdapterTree(adapters, len(adapters))
	return getAllAdapterTreeLeaves(tree[0], tree, make([]int, len(tree)))
}

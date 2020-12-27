package main

func getMultipliedJoltDifferences(adapters []int) int {
	joltDifferences := map[int]int{1: 1, 3: 1}
	for i := 1; i < len(adapters); i++ {
		difference := adapters[i] - adapters[i-1]
		joltDifferences[difference]++
	}
	return joltDifferences[1] * joltDifferences[3]
}

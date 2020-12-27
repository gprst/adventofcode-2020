package main

import "sort"

func sum(numbers ...int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func getWeaknessContiguousRange(inputs []int, weaknessNumber int) (contiguousRange []int) {
	index := 0
	for sum(contiguousRange...) != weaknessNumber {
		// Add next number
		contiguousRange = append(contiguousRange, inputs[index])
		// Remove first numbers if sum is too high
		for sum(contiguousRange...) > weaknessNumber {
			contiguousRange = contiguousRange[1:]
		}
		index++
	}
	return contiguousRange
}

func getEncryptionWeakness(inputs []int) int {
	weaknessNumber := getWeaknessNumber(inputs)
	contiguousRange := getWeaknessContiguousRange(inputs, weaknessNumber)
	sort.Ints(contiguousRange)
	return contiguousRange[0] + contiguousRange[len(contiguousRange)-1]
}

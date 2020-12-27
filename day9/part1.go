package main

func isOK(number int, preamble []int, preambleSize int) bool {
	for firstIndex := 0; firstIndex < preambleSize-1; firstIndex++ {
		for secondIndex := firstIndex + 1; secondIndex < preambleSize; secondIndex++ {
			if preamble[firstIndex]+preamble[secondIndex] == number {
				return true
			}
		}
	}
	return false
}

func getWeaknessNumber(inputs []int) int {
	const preambleSize = 25
	preamble := inputs[:preambleSize]
	index := preambleSize
	for isOK(inputs[index], preamble, preambleSize) {
		index++
		preamble = inputs[index-preambleSize : index]
	}
	return inputs[index]
}

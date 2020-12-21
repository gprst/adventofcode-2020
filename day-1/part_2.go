package main

func multiplyThreeNumbersThatSumUpTo2020(numbers []int) int {
	const goal = 2020
	for bigInputIndex := len(numbers) - 1; bigInputIndex >= 0; bigInputIndex-- {
		bigInput := numbers[bigInputIndex]
		for smallInputIndex, smallInput := range numbers {
			if bigInput+smallInput > goal {
				break
			} else {
				for lastInputIndex := smallInputIndex; lastInputIndex < bigInputIndex; lastInputIndex++ {
					lastInput := numbers[lastInputIndex]
					if bigInput+lastInput+smallInput > goal {
						break
					} else if bigInput+smallInput+lastInput == goal {
						return bigInput * smallInput * lastInput
					}
				}
			}
		}
	}
	return 0
}

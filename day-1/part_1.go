package day1

func multiplyTwoNumbersThatSumUpTo2020(numbers []int) int {
	const goal = 2020
	for i := len(numbers) - 1; i >= 0; i-- {
		bigInput := numbers[i]
		for _, smallInput := range numbers {
			if bigInput+smallInput == goal {
				return smallInput * bigInput
			} else if bigInput+smallInput > goal {
				break
			}
		}
	}
	return 0
}

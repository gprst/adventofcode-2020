package day1

import (
	"fmt"
	"sort"
)

func main() {
	inputs, err := scanLines("./input")
	if err != nil {
		panic(err)
	}

	sort.Ints(inputs)

	fmt.Println("The answer for the first part is:", multiplyTwoNumbersThatSumUpTo2020(inputs))
	fmt.Println("The answer for the second part is:", multiplyThreeNumbersThatSumUpTo2020(inputs))
}

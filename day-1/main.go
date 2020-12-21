package main

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

	fmt.Println("The answer is:", multiplyTwoNumbersThatSumUpTo2020(inputs))
}

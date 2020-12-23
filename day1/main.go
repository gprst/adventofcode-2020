package main

import (
	"adventofcode/2020/linescanner"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	inputs, err := linescanner.ScanLines(curDir + "/input")
	if err != nil {
		panic(err)
	}

	var numbers []int

	for _, input := range inputs {
		number, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	sort.Ints(numbers)

	fmt.Println("The answer for the first part is:", multiplyTwoNumbersThatSumUpTo2020(numbers))
	fmt.Println("The answer for the second part is:", multiplyThreeNumbersThatSumUpTo2020(numbers))
}

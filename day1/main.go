package main

import (
	"adventofcode/2020/linescanner"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

	sort.Ints(inputs)

	fmt.Println("The answer for the first part is:", multiplyTwoNumbersThatSumUpTo2020(inputs))
	fmt.Println("The answer for the second part is:", multiplyThreeNumbersThatSumUpTo2020(inputs))
}

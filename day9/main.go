package main

import (
	"adventofcode/2020/inputscanner"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	rawInputs, err := inputscanner.ScanLines(curDir + "/input")
	if err != nil {
		panic(err)
	}

	var inputs []int
	for _, rawInput := range rawInputs {
		number, _ := strconv.Atoi(rawInput)
		inputs = append(inputs, number)
	}

	fmt.Println("The answer for the first part is:", getWeaknessNumber(inputs))
}

package main

import (
	"adventofcode/2020/linescanner"
	"fmt"
	"os"
	"path/filepath"
)

var patternLength uint16

func main() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	inputs, err := linescanner.ScanLines(curDir + "/input")
	if err != nil {
		panic(err)
	}

	patternLength = uint16(len(inputs[0]))

	fmt.Println("The answer for the first part is:", getEncounteredTreesCount(inputs))
	fmt.Println("The answer for the second part is:", getMultipliedEncounteredTreesCount(inputs))
}

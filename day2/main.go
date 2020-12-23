package main

import (
	"adventofcode/2020/linescanner"
	"fmt"
	"os"
	"path/filepath"
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

	fmt.Println("The answer for the first part is:", checkPasswordDatabase(inputs))
}

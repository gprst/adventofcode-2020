package linescanner

import (
	"bufio"
	"os"
	"strconv"
)

// ScanLines returns an array of integers containing all lines of the file whose
// path is given as input
func ScanLines(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []int

	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, input)
	}

	return lines, nil
}

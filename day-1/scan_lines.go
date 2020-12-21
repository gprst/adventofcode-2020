package main

import (
	"bufio"
	"os"
	"strconv"
)

func scanLines(filePath string) ([]int, error) {
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

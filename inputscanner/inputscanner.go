package inputscanner

import (
	"bufio"
	"io/ioutil"
	"os"
)

// ScanLines returns an array of strings containing all lines of the file whose
// path is given as input
func ScanLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// ScanInput returns the content of the file whose path is given as argument as
// a string
func ScanInput(filepath string) (string, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

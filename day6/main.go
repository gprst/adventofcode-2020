package main

import (
	"adventofcode/2020/inputscanner"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getGroupsAnswers(input string) (groupsAnswers [][]string) {
	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		groupsAnswers = append(groupsAnswers, strings.Fields(group))
	}
	return groupsAnswers
}

func main() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	input, err := inputscanner.ScanInput(curDir + "/input")
	if err != nil {
		panic(err)
	}

	groupsAnswers := getGroupsAnswers(input)

	fmt.Println("The answer for the first part is:", sumUpGroupsPositiveAnswers(groupsAnswers))
	fmt.Println("The answer for the first part is:", sumUpGroupsUnanimousAnswers(groupsAnswers))
}

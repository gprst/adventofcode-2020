package main

import (
	"strconv"
	"strings"
)

func getKeyNumbers(rawKeyNumbers string) (uint8, uint8) {
	splitKeyNumbers := strings.Split(rawKeyNumbers, "-")
	firstKeyNumber, _ := strconv.Atoi(splitKeyNumbers[0])
	secondKeyNumber, _ := strconv.Atoi(splitKeyNumbers[1])
	return uint8(firstKeyNumber), uint8(secondKeyNumber)
}

func getMandatoryLetter(input string) rune {
	return rune(input[0])
}

func parseInput(rawSplitInput []string) (firstKeyNumber uint8, secondKeyNumber uint8, mandatoryLetter rune, password string) {
	firstKeyNumber, secondKeyNumber = getKeyNumbers(rawSplitInput[0])
	mandatoryLetter = rune(rawSplitInput[1][0])
	password = rawSplitInput[2]
	return firstKeyNumber, secondKeyNumber, mandatoryLetter, password
}

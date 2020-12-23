package main

import (
	"strconv"
	"strings"
)

func getMinAndMax(minAndMax string) (uint8, uint8) {
	splitMinAndMax := strings.Split(minAndMax, "-")
	min, _ := strconv.Atoi(splitMinAndMax[0])
	max, _ := strconv.Atoi(splitMinAndMax[1])
	return uint8(min), uint8(max)
}

func getMandatoryLetter(input string) rune {
	return rune(input[0])
}

func parseInput(rawSplitInput []string) (min uint8, max uint8, mandatoryLetter rune, password string) {
	min, max = getMinAndMax(rawSplitInput[0])
	mandatoryLetter = rune(rawSplitInput[1][0])
	password = rawSplitInput[2]
	return min, max, mandatoryLetter, password
}

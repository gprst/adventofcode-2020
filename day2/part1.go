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

func getMandatoryLetterCount(mandatoryLetter rune, password string) uint8 {
	var count uint8
	for _, letter := range password {
		if letter == mandatoryLetter {
			count++
		}
	}
	return count
}

func checkPasswordDatabase(inputs []string) uint16 {
	var validPasswords uint16
	for _, input := range inputs {
		rawSplitInput := strings.Split(input, " ")
		min, max := getMinAndMax(rawSplitInput[0])
		mandatoryLetter := rune(rawSplitInput[1][0])
		password := rawSplitInput[2]
		mandatoryLetterCount := getMandatoryLetterCount(mandatoryLetter, password)
		if mandatoryLetterCount >= min && mandatoryLetterCount <= max {
			validPasswords++
		}
	}
	return validPasswords
}

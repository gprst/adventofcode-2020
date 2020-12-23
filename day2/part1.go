package main

import "strings"

func getMandatoryLetterCount(mandatoryLetter rune, password string) uint8 {
	var count uint8
	for _, letter := range password {
		if letter == mandatoryLetter {
			count++
		}
	}
	return count
}

func checkPasswordDatabase(inputs []string) (validPasswords uint16) {
	for _, input := range inputs {
		rawSplitInput := strings.Split(input, " ")
		min, max, mandatoryLetter, password := parseInput(rawSplitInput)
		mandatoryLetterCount := getMandatoryLetterCount(mandatoryLetter, password)
		if mandatoryLetterCount >= min && mandatoryLetterCount <= max {
			validPasswords++
		}
	}
	return validPasswords
}

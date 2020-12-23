package main

import "strings"

func hasIndexMandatoryLetter(index uint8, word string, mandatoryLetter rune) bool {
	if word[index-1] == byte(mandatoryLetter) {
		return true
	}
	return false
}

func checkPasswordDatabaseWithNewPolicy(inputs []string) (validPasswords uint16) {
	for _, input := range inputs {
		var numberOfWellPlacedMandatoryLetters uint8
		rawSplitInput := strings.Split(input, " ")
		firstKeyNumber, secondKeyNumber, mandatoryLetter, password := parseInput(rawSplitInput)
		if hasIndexMandatoryLetter(firstKeyNumber, password, mandatoryLetter) {
			numberOfWellPlacedMandatoryLetters++
		}
		if hasIndexMandatoryLetter(secondKeyNumber, password, mandatoryLetter) {
			numberOfWellPlacedMandatoryLetters++
		}
		if numberOfWellPlacedMandatoryLetters == 1 {
			validPasswords++
		}
	}
	return validPasswords

}

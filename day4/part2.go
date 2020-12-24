package main

import (
	"regexp"
	"strconv"
)

func isYearValid(value string, min int, max int) bool {
	year, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return year >= min && year <= max
}

func isPatternValid(value string, pattern string) bool {
	isValid, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false
	}
	return isValid
}

func isHgtValid(height string) bool {
	splitIndex := len(height) - 2
	unit := height[splitIndex:]
	value, err := strconv.Atoi(height[:splitIndex])
	if err != nil {
		return false
	}
	if unit == "cm" {
		return value >= 150 && value <= 193
	} else if unit == "in" {
		return value >= 59 && value <= 76
	}
	return false
}

func isEclValid(eyeColor string) bool {
	validEyeColors := [...]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validEyeColor := range validEyeColors {
		if eyeColor == validEyeColor {
			return true
		}
	}
	return false
}

func isFieldValid(field string, value string) bool {
	switch field {
	case "byr":
		return isYearValid(value, 1920, 2002)
	case "iyr":
		return isYearValid(value, 2010, 2020)
	case "eyr":
		return isYearValid(value, 2020, 2030)
	case "hgt":
		return isHgtValid(value)
	case "hcl":
		return isPatternValid(value, `^#[a-f0-9]{6}$`)
	case "ecl":
		return isEclValid(value)
	case "pid":
		return isPatternValid(value, `^\d{9}$`)
	case "cid":
		return true
	default:
		return false
	}
}

func howManyDocumentsWithValidFields(documents []map[string]string) (validDocuments uint) {
	for _, document := range documents {
		if hasDocumentRequiredFields(document) {
			hasDocumentValidFields := true
			for field, value := range document {
				if !isFieldValid(field, value) {
					hasDocumentValidFields = false
					break
				}
			}
			if hasDocumentValidFields {
				validDocuments++
			}
		}
	}
	return validDocuments
}

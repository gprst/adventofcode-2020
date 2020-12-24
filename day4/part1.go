package main

func hasCid(document map[string]string) bool {
	for property := range document {
		if property == "cid" {
			return true
		}
	}
	return false
}

func isDocumentValid(document map[string]string) bool {
	validProperties := [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	isValidWithCid := hasCid(document) && len(document) == len(validProperties)+1
	isValidWithoutCid := !hasCid(document) && len(document) == len(validProperties)
	if isValidWithCid || isValidWithoutCid {
		return true
	}
	return false
}

func howManyValidDocuments(documents []map[string]string) (validDocuments uint) {
	for _, document := range documents {
		if isDocumentValid(document) {
			validDocuments++
		}
	}
	return validDocuments
}

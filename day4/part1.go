package main

func howManyValidDocuments(documents []map[string]string) (validDocuments uint) {
	for _, document := range documents {
		if hasDocumentRequiredFields(document) {
			validDocuments++
		}
	}
	return validDocuments
}

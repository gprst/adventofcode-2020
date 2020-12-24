package main

import (
	"adventofcode/2020/inputscanner"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func mapDocumentProperties(properties []string) map[string]string {
	mappedDocumentProperties := make(map[string]string)
	for _, property := range properties {
		keyValue := strings.Split(property, ":")
		mappedDocumentProperties[keyValue[0]] = keyValue[1]
	}
	return mappedDocumentProperties
}

func mapDocuments(input string) (mappedDocuments []map[string]string) {
	rawDocuments := strings.Split(input, "\n\n")
	for _, document := range rawDocuments {
		documentProperties := strings.Fields(document)
		mappedDocuments = append(mappedDocuments, mapDocumentProperties(documentProperties))
	}
	return mappedDocuments
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

	mappedDocuments := mapDocuments(input)

	fmt.Println("The answer for the first part is:", howManyValidDocuments(mappedDocuments))
	fmt.Println("The answer for the first part is:", howManyDocumentsWithValidFields(mappedDocuments))
}

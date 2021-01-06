package main

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	mask = "mask"
	mem  = "mem"
)

func overwriteValue(value int, mask []byte) int {
	binaryValue := strconv.FormatInt(int64(value), 2)
	byteValue := []byte(strings.Repeat("0", 36-len(binaryValue)) + binaryValue)
	for i := range byteValue {
		maskByte := mask[i]
		if maskByte == '0' || maskByte == '1' {
			byteValue[i] = maskByte
		}
	}
	overwrittenValue, _ := strconv.ParseInt(string(byteValue), 10, 36)
	return int(overwrittenValue)
}

func parseMemOperation(operation []string) (int, int) {
	value, _ := strconv.Atoi(operation[1])

	targetPattern := regexp.MustCompile(`\[(\d+)\]`)
	target, _ := strconv.Atoi(targetPattern.FindStringSubmatch(operation[0])[1])

	return target, value
}

func sumValuesInMemory(operations []string) int {
	mem := make(map[int]int)

	var curMask string
	for _, operation := range operations {
		splitOperation := strings.Split(operation, " = ")

		if splitOperation[0][:4] == mask {
			curMask := splitOperation[1]
			continue
		}

		target, value := parseMemOperation(splitOperation)

		overwrittenValue := overwriteValue(value, []byte(curMask))
	}

	return 0
}

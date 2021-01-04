package main

import (
	"math"
	"strconv"
)

type state struct {
	x, y      int
	direction string
}

var directions = []string{"N", "E", "S", "W"}

func getDirectionIndex(curDirection string) int {
	for index, direction := range directions {
		if direction == curDirection {
			return index
		}
	}
	panic("Direction doesn't exist")
}

func updateDirection(curDirection *string, instructionType string, instructionValue int) {
	steps := instructionValue / 90
	directionIndex := getDirectionIndex(*curDirection)

	if instructionType == "R" {
		directionIndex += steps
	} else if instructionType == "L" {
		directionIndex -= steps
	}

	if directionIndex >= len(directions) {
		directionIndex -= len(directions)
	} else if directionIndex < 0 {
		directionIndex += len(directions)
	}

	*curDirection = directions[directionIndex]
}

func updateState(shipState *state, instruction string) {
	instructionType := instruction[:1]
	instructionValue, _ := strconv.Atoi(instruction[1:])

	if instructionType == "F" {
		instructionType = shipState.direction
	}

	switch instructionType {
	case "N":
		shipState.y += instructionValue
		break
	case "S":
		shipState.y -= instructionValue
		break
	case "E":
		shipState.x += instructionValue
		break
	case "W":
		shipState.x -= instructionValue
		break
	default:
		updateDirection(&shipState.direction, instructionType, instructionValue)
	}
}

func getManhattanDistance(inputs []string) int {
	shipState := state{0, 0, "E"}

	for _, instruction := range inputs {
		updateState(&shipState, instruction)
	}

	return int(math.Abs(float64(shipState.x)) + math.Abs(float64(shipState.y)))
}

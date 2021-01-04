package main

import (
	"math"
	"strconv"
)

type coordinates struct {
	x, y int
}

const (
	north   = "N"
	south   = "S"
	east    = "E"
	west    = "W"
	left    = "L"
	right   = "R"
	forward = "F"
)

func rotateWaypoint(waypoint *coordinates, instructionType string, instructionValue int) {
	steps := instructionValue / 90
	if instructionType == right {
		for i := 0; i < steps; i++ {
			y := waypoint.y
			waypoint.y = -waypoint.x
			waypoint.x = y
		}
	} else if instructionType == left {
		for i := 0; i < steps; i++ {
			x := waypoint.x
			waypoint.x = -waypoint.y
			waypoint.y = x
		}
	}
}

func moveWaypoint(waypoint *coordinates, instructionType string, instructionValue int) {
	if instructionType == left || instructionType == right {
		rotateWaypoint(waypoint, instructionType, instructionValue)
	} else {
		switch instructionType {
		case north:
			waypoint.y += instructionValue
			break
		case south:
			waypoint.y -= instructionValue
			break
		case east:
			waypoint.x += instructionValue
			break
		case west:
			waypoint.x -= instructionValue
		}
	}
}

func moveShip(shipPosition *coordinates, waypoint coordinates, instructionValue int) {
	shipPosition.x += waypoint.x * instructionValue
	shipPosition.y += waypoint.y * instructionValue
}

func getManhattanDistanceWithWaypoint(inputs []string) int {
	shipPosition := coordinates{0, 0}
	waypoint := coordinates{10, 1}

	for _, instruction := range inputs {
		instructionType := instruction[:1]
		instructionValue, _ := strconv.Atoi(instruction[1:])
		if instructionType == forward {
			moveShip(&shipPosition, waypoint, instructionValue)
		} else {
			moveWaypoint(&waypoint, instructionType, instructionValue)
		}
	}

	return int(math.Abs(float64(shipPosition.x)) + math.Abs(float64(shipPosition.y)))
}

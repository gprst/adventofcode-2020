package main

import (
	"strconv"
	"strings"
)

func getBusToTakeAndDeparture(firstBus int, departures map[int]int) (int, int) {
	busToTake := firstBus
	earliestDeparture := departures[firstBus]
	for busID, departureTimestamp := range departures {
		if departureTimestamp < earliestDeparture {
			earliestDeparture = departureTimestamp
			busToTake = busID
		}
	}
	return busToTake, earliestDeparture
}

func getBusesEarliestDepartures(buses []int, airportArrival int) map[int]int {
	departures := make(map[int]int, len(buses))
	for _, busLoopTime := range buses {
		departureTimestamp := 0
		for departureTimestamp < airportArrival {
			departureTimestamp += busLoopTime
		}
		departures[busLoopTime] = departureTimestamp
	}
	return departures
}

func firstSolution(inputs []string) int {
	airportArrival, _ := strconv.Atoi(inputs[0])
	busesIDs := strings.Split(strings.ReplaceAll(inputs[1], "x,", ""), ",")

	buses := make([]int, len(busesIDs))
	for index, busID := range busesIDs {
		buses[index], _ = strconv.Atoi(busID)
	}

	departures := getBusesEarliestDepartures(buses, airportArrival)

	busToTake, earliestDeparture := getBusToTakeAndDeparture(buses[0], departures)

	return (earliestDeparture - airportArrival) * busToTake
}

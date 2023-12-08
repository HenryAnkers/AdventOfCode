package main

import (
	"fmt"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day8.txt"
	input, err := utils.ReadLines(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start := time.Now()
	partOne(input)
	duration1 := time.Since(start)

	start = time.Now()
	partTwo(input)
	duration2 := time.Since(start)
	fmt.Printf("Part One took: %v, Part Two took: %v\n", duration1, duration2)
}

func partOne(input []string) {
	numTurns := 0
	turnSequence := input[0]
	nextTurnIndex := map[string]int{"L": 0, "R": 1}
	networkMap := buildNetworkMap(input[2:])
	location := "AAA"

	for location != "ZZZ" {
		nextTurn := turnSequence[numTurns%len(turnSequence)]
		location = networkMap[location][nextTurnIndex[string(nextTurn)]]
		numTurns++
	}

	fmt.Println("Part One:", numTurns)
}

func partTwo(input []string) {
	// brute force didn't work
	// as there are matching nodes ending with A as nodes ending with Z, hopefully cycles are all disjoint and only involve one A and one Z node
	// therefore can find LCM of the numbers

	answer := 1
	turnSequence := input[0]
	nextTurnIndex := map[string]int{"L": 0, "R": 1}
	networkMap := buildNetworkMap(input[2:])
	locations := []string{}

	for location, _ := range networkMap {
		if string(location[2]) == "A" {
			locations = append(locations, location)
		}
	}

	cycleLengths := []int{}
	for _, location := range locations {
		currentLoc := location
		length := 0

		for string(currentLoc[2]) != "Z" {
			nextTurn := turnSequence[length%len(turnSequence)]
			currentLoc = networkMap[currentLoc][nextTurnIndex[string(nextTurn)]]
			length++
		}

		cycleLengths = append(cycleLengths, length)
	}

	for _, length := range cycleLengths {
		answer = utils.LCM(answer, length)
	}

	fmt.Println("Part Two:", answer)
}

func buildNetworkMap(input []string) map[string][]string {
	mapOutput := map[string][]string{}

	for _, line := range input {
		location := strings.Split(line, " = ")[0]
		lAndRLocString := strings.Split(line, " = ")[1]
		leftLoc := strings.Split(lAndRLocString, ", ")[0][1:]
		rightLoc := strings.Split(lAndRLocString, ", ")[1][:3]

		mapOutput[location] = []string{leftLoc, rightLoc}
	}

	return mapOutput
}

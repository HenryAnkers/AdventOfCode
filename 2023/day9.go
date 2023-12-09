package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day9.txt"
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
	answer := 0

	for _, line := range input {
		numberDiffs := calculateNumberDiffs(line)

		for i := len(numberDiffs) - 2; i >= 0; i-- {
			previousDelta := numberDiffs[i+1][len(numberDiffs[i+1])-1]
			nextValue := numberDiffs[i][len(numberDiffs[i])-1] + previousDelta
			numberDiffs[i] = append(numberDiffs[i], nextValue)

			if i == 0 {
				answer += nextValue
			}
		}
	}

	fmt.Println("Part One:", answer)
}

func partTwo(input []string) {
	answer := 0

	for _, line := range input {
		numberDiffs := calculateNumberDiffs(line)

		for i := len(numberDiffs) - 2; i >= 0; i-- {
			//we now want to add to the beginning of the list instead, so take delta away from index[0] and add to beginning of array
			previousDelta := numberDiffs[i+1][0]
			nextValue := numberDiffs[i][0] - previousDelta
			numberDiffs[i] = append([]int{nextValue}, numberDiffs[i]...)

			if i == 0 {
				answer += nextValue
			}
		}
	}

	fmt.Println("Part Two:", answer)
}

func calculateNumberDiffs(line string) [][]int {
	numberDiffs := [][]int{}

	currentDiff := []int{}
	for _, strInt := range strings.Split(line, " ") {
		parsedInt, _ := strconv.Atoi(string(strInt))
		currentDiff = append(currentDiff, parsedInt)
	}
	numberDiffs = append(numberDiffs, currentDiff)

	for true {
		newDiff := []int{}

		allZero := true
		for i := 1; i < len(currentDiff); i++ {
			nextDiffValue := currentDiff[i] - currentDiff[i-1]
			if nextDiffValue != 0 {
				allZero = false
			}
			newDiff = append(newDiff, nextDiffValue)
		}

		currentDiff = newDiff
		numberDiffs = append(numberDiffs, currentDiff)

		if allZero {
			break
		}
	}
	return numberDiffs
}

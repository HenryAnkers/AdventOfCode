package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day2.txt"
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
	total := 0
	maxValues := map[string]int{"red": 12, "green": 13, "blue": 14}

games:
	for i, gameString := range input {
		roundsPlayed := parseRounds(gameString)

		for _, round := range roundsPlayed {
			for key, maxAllowed := range maxValues {
				if round[key] > maxAllowed {
					continue games // this round is invalid, so move on to the next game
				}
			}
		}

		total += (i + 1)
	}

	fmt.Printf("Part One: %d\n", total)
}

func partTwo(input []string) {
	total := 0

	for _, gameString := range input {
		maxFound := map[string]int{"red": 0, "green": 0, "blue": 0}
		roundsPlayed := parseRounds(gameString)
		powerOfSet := 1

		for _, round := range roundsPlayed {
			for key, currentMax := range maxFound {
				maxFound[key] = max(currentMax, round[key]) //the maximum amount of any colour we see in this round is the minimum we require
			}
		}

		for _, amount := range maxFound {
			if amount != 0 {
				powerOfSet *= amount
			}
		}

		total += powerOfSet
	}

	fmt.Printf("Part Two: %d\n", total)
}

func parseRounds(input string) []map[string]int {
	roundsString := strings.Split(input, ": ")[1]
	rounds := strings.Split(roundsString, "; ")
	results := []map[string]int{}

	for _, round := range rounds {
		values := map[string]int{"red": 0, "green": 0, "blue": 0}

		pulls := strings.Split(round, ", ")
		for _, pull := range pulls {
			amount, _ := strconv.Atoi(strings.Split(pull, " ")[0])
			colour := strings.Split(pull, " ")[1]
			values[colour] += amount
		}

		results = append(results, values)
	}

	return results
}

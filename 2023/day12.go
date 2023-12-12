package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day12.txt"
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
	// at every ? we can change to "." or "#" and see if the remaining string is valid or not
	// we can cache the states of [remainingString + "-" + remainingSpringsToFind], do we even need this?
	// not sure this is good enough but lets see

	answer := 0

	for _, line := range input {
		recordList := strings.Split(line, " ")[0]
		springStrList := strings.Split(strings.Split(line, " ")[1], ",")
		springList := []int{}
		for _, s := range springStrList {
			digit, _ := strconv.Atoi(string(s))
			springList = append(springList, digit)
		}

		curr := getRecordArrangements(recordList, springList, 0, map[string]int{})
		answer += curr
	}

	fmt.Println("Part One:", answer)
}

func partTwo(input []string) {
	// I think we just need to memo states
	answer := 0

	for _, line := range input {
		recordList := strings.Split(line, " ")[0]
		springStrList := strings.Split(line, " ")[1]

		bigRecordList := recordList
		bigSpringStr := springStrList
		for i := 0; i < 4; i++ {
			bigRecordList += ("?" + recordList)
			bigSpringStr += ("," + springStrList)
		}
		bigSpringStrList := strings.Split(bigSpringStr, ",")

		springList := []int{}
		for _, s := range bigSpringStrList {
			digit, _ := strconv.Atoi(string(s))
			springList = append(springList, digit)
		}

		curr := getRecordArrangements(bigRecordList, springList, 0, map[string]int{})
		answer += curr
	}

	fmt.Println("Part Two:", answer)
}

func getRecordArrangements(record string, springsToFind []int, currentSpringLength int, memo map[string]int) int {
	key := getCacheKey(record, springsToFind, currentSpringLength)
	if prevAns, ok := memo[key]; ok {
		return prevAns
	}

	if record == "" {
		if len(springsToFind) == 0 && currentSpringLength == 0 || (len(springsToFind) == 1 && currentSpringLength == springsToFind[0]) {
			return 1
		}
		return 0
	}

	if record[0] == '#' {
		ans := getRecordArrangements(record[1:], springsToFind, currentSpringLength+1, memo)
		memo[key] = ans
		return ans
	}

	if record[0] == '.' {
		if currentSpringLength > 0 {
			// we're at the end of the spring....
			if len(springsToFind) == 0 || springsToFind[0] != currentSpringLength {
				// either we dont want any more springs or this one has wrong length
				memo[key] = 0
				return 0
			} else {
				ans := getRecordArrangements(record[1:], springsToFind[1:], 0, memo)
				memo[key] = ans
				return ans
			}
		} else {
			ans := getRecordArrangements(record[1:], springsToFind, 0, memo)
			memo[key] = ans
			return ans
		}
	}

	if record[0] == '?' {
		ans := getRecordArrangements(record[1:], springsToFind, currentSpringLength+1, memo) + getRecordArrangements("."+record[1:], springsToFind, currentSpringLength, memo)
		memo[key] = ans
		return ans
	}

	fmt.Println("invalid input or something")
	return 0
}

func getCacheKey(record string, springsToFind []int, currentSpringLength int) string {
	key := record + " -- "
	for _, spring := range springsToFind {
		key += (strconv.Itoa(spring) + ",")
	}
	return key + " -- " + strconv.Itoa(currentSpringLength)
}

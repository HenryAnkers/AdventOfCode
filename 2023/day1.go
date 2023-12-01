package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day1.txt"
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
	for _, s := range input {
		var firstInt, lastInt string

		for _, char := range s {
			if unicode.IsDigit(char) {
				if firstInt == "" {
					firstInt = string(char)
				}
				lastInt = string(char)
			}
		}

		localTotal, _ := strconv.Atoi(firstInt + lastInt)
		total += localTotal
	}

	fmt.Printf("Part One: %d\n", total)
}

func partTwo(input []string) {
	total := 0
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, s := range input {
		var firstInt, lastInt string
		pointer := 0

	search:
		for pointer < len(s) {
			for k, v := range numberMap {
				if strings.HasPrefix(s[pointer:], k) {
					if firstInt == "" {
						firstInt = v
					}
					lastInt = v
					pointer += len(k) - 1 // -1 because of strings like twone which need to have firstInt=2, lastInt=1. this is also why we can't just do strings.Replace
					continue search
				}
			}

			if unicode.IsDigit(rune(s[pointer])) {
				if firstInt == "" {
					firstInt = string(s[pointer])
				}
				lastInt = string(s[pointer])
			}

			pointer++
		}

		localTotal, _ := strconv.Atoi(firstInt + lastInt)
		total += localTotal
	}

	fmt.Printf("Part Two: %d\n", total)
}

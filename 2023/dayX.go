package main

import (
	"fmt"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/dayX.txt"
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

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	answer := 0

	fmt.Printf("Part Two: %d\n", answer)
}

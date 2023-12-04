package main

import (
	u "adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFilePath := "./input/day4.txt"
	input, err := u.ReadLines(inputFilePath)
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

	for _, row := range input {
		localAns := 0
		card, winningNumbers := parseCardRow(row)

		for _, number := range card {
			if winningNumbers[number] {
				if localAns != 0 {
					localAns *= 2
				} else {
					localAns = 1
				}
			}
		}

		answer += localAns
	}

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	answer := 0

	numberOfGames := map[int]int{}
	for k, _ := range input {
		numberOfGames[k+1] = 1
	}

	for index, row := range input {
		gameNumber := index + 1
		totalMatches := 0
		numberOfInstancesOfGame := numberOfGames[gameNumber]

		card, winningNumbers := parseCardRow(row)
		for _, number := range card {
			if winningNumbers[number] {
				totalMatches += 1
			}
		}

		for i := 1; i <= totalMatches; i++ {
			gameNumberWon := gameNumber + i
			numberOfGames[gameNumberWon] += numberOfInstancesOfGame
		}
	}

	for _, v := range numberOfGames {
		answer += v
	}

	fmt.Printf("Part Two: %d\n", answer)
}

func parseCardRow(input string) ([]int, map[int]bool) {
	cardString := strings.Split(input, " | ")[0]
	cardNumbersString := strings.Split(cardString, ": ")[1]
	winningNumbersString := strings.Split(input, " | ")[1]
	cardNumbersStringList := strings.Split(cardNumbersString, " ")
	winningNumbersStringList := strings.Split(winningNumbersString, " ")

	cardNumbers := []int{}
	winningNumbers := map[int]bool{}

	for _, value := range cardNumbersStringList {
		digit, err := strconv.Atoi(value)

		if err == nil {
			cardNumbers = append(cardNumbers, digit)
		}
	}

	for _, value := range winningNumbersStringList {
		digit, err := strconv.Atoi(value)

		if err == nil {
			winningNumbers[digit] = true
		}
	}

	return cardNumbers, winningNumbers
}

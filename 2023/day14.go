package main

import (
	"fmt"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day14.txt"
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

	inputMap := map[int]map[int]rune{}
	for y, line := range input {
		inputMap[y] = map[int]rune{}
		for x, char := range line {
			inputMap[y][x] = char
		}
	}

	for y := 1; y < len(inputMap); y++ {
		for x := 0; x < len(inputMap[y]); x++ {
			char := inputMap[y][x]
			if char == 'O' {
				newY := y
				for i := 1; i <= y; i++ {
					if inputMap[y-i][x] == '.' {
						newY = y - i
					} else {
						break
					}
				}
				if y != newY {
					inputMap[newY][x] = 'O'
					inputMap[y][x] = '.'
				}
			}
		}
	}

	for y, _ := range inputMap {
		for _, char := range inputMap[y] {
			if char == 'O' {
				answer += (len(inputMap)) - y
			}
		}
	}

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// probably just save a string map of the state after each cycle, wait until we see a pattern, then the pattern length L is last-first...
	// we can then skip floor((1000000000 - number of steps already processed) / L) * L steps
	answer := 0

	inputMap := map[int]map[int]rune{}
	for y, line := range input {
		inputMap[y] = map[int]rune{}
		for x, char := range line {
			inputMap[y][x] = char
		}
	}

	stateCache := map[string]int{}

	for X := 0; X < 1000000000; X++ {
		if prevFound, ok := stateCache[getStateKey(inputMap)]; ok {
			skip := ((1000000000 - X) / (X - prevFound)) * (X - prevFound)
			X += skip
		} else {
			stateCache[getStateKey(inputMap)] = X
		}

		// North
		for y := 1; y < len(inputMap); y++ {
			for x := 0; x < len(inputMap[y]); x++ {
				char := inputMap[y][x]
				if char == 'O' {
					newY := y
					for i := 1; i <= y; i++ {
						if inputMap[y-i][x] == '.' {
							newY = y - i
						} else {
							break
						}
					}
					if y != newY {
						inputMap[newY][x] = 'O'
						inputMap[y][x] = '.'
					}
				}
			}
		}

		// West
		for y := 0; y < len(inputMap); y++ {
			for x := 0; x < len(inputMap[y]); x++ {
				char := inputMap[y][x]
				if char == 'O' {
					newX := x
					for i := 1; i <= x; i++ {
						if inputMap[y][x-i] == '.' {
							newX = x - i
						} else {
							break
						}
					}
					if x != newX {
						inputMap[y][newX] = 'O'
						inputMap[y][x] = '.'
					}
				}
			}
		}

		//South
		for y := len(inputMap) - 1; y >= 0; y-- {
			for x := 0; x < len(inputMap[y]); x++ {
				char := inputMap[y][x]
				if char == 'O' {
					newY := y
					for i := 1; i+y < len(inputMap); i++ {
						if inputMap[y+i][x] == '.' {
							newY = y + i
						} else {
							break
						}
					}

					if y != newY {
						inputMap[newY][x] = 'O'
						inputMap[y][x] = '.'
					}
				}
			}
		}

		// East
		for y := 0; y < len(inputMap); y++ {
			for x := len(inputMap[y]) - 1; x >= 0; x-- {
				char := inputMap[y][x]
				if char == 'O' {
					newX := x
					for i := 1; i+x < len(inputMap[y]); i++ {
						if inputMap[y][x+i] == '.' {
							newX = x + i
						} else {
							break
						}
					}
					if x != newX {
						inputMap[y][newX] = 'O'
						inputMap[y][x] = '.'
					}
				}
			}
		}
	}

	for y, _ := range inputMap {
		for _, char := range inputMap[y] {
			if char == 'O' {
				answer += (len(inputMap)) - y
			}
		}
	}

	fmt.Printf("Part Two: %d\n", answer)
}

func getStateKey(inputMap map[int]map[int]rune) string {
	key := ""
	for y := 0; y < len(inputMap); y++ {
		str := ""
		for x := 0; x < len(inputMap[y]); x++ {
			str += string(inputMap[y][x])
		}
		key += " - " + str
	}
	return key
}

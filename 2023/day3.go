package main

import (
	"fmt"
	"strconv"
	"time"
	"unicode"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day3.txt"
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
	// lets scan through the input l>r, row by row.
	// at every point if we see a '.' we continue, if we see a part we also continue for now
	// once we see a digit (unicode.IsDigit(char))
	//  - create a string to hold the digits
	//  - check all 9 squares around the point to see if there is a part character (lets say not a . and not a digit), if so mark boolean isPartNumber true
	//  - check the next right digit if possible to see if it is a digit, if so repeat process,
	//  - if not add the number to total iff it's a part number, reset the digit string and boolean, and continue on

	answer := 0

	for y, row := range input {
		currentDigitString := ""
		isPartNumber := false

		for x, char := range row {
			// handle building digit string
			if currentDigitString != "" {
				if unicode.IsDigit(char) {
					if !isPartNumber {
						isPartNumber, _, _, _ = checkNearPartNumber(input, x, y)
					}
					currentDigitString += string(char)
				}

				// we've found the end of the digit (or line), convert the current digit string to a int and then reset all variables
				if !unicode.IsDigit(char) || x == len(row)-1 {
					if isPartNumber {
						digit, _ := strconv.Atoi(currentDigitString)
						answer += digit
					}
					currentDigitString = ""
					isPartNumber = false
				}
			} else {
				// handle outside of digit string
				if unicode.IsDigit(char) {
					currentDigitString += string(char)
					isPartNumber, _, _, _ = checkNearPartNumber(input, x, y)
				}
			}
		}
	}

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// on top of our part 1 answer, save a map[int]map[int][]int where each key (x,y) represents a coordinate of a "*" character and each value represents the list of adjacent part numbers
	// we cycle through this at the end and add every one with len(2) to the answer
	// modify checkIsPartNumber to also return the part

	answer := 0
	gearPartNumberLists := map[u.Coord][]int{}

	for y, row := range input {
		currentDigitString := ""
		gearsFoundForDigit := map[u.Coord]bool{}

		for x, char := range row {
			// handle building digit string
			if currentDigitString != "" {
				if unicode.IsDigit(char) {
					_, part, partX, partY := checkNearPartNumber(input, x, y)

					// if we find a gear, add it to the list. we can't just add to gearPartNumberLists as we don't know whole number yet
					if part == "*" {
						gearsFoundForDigit[u.Coord{X: partX, Y: partY}] = true
					}

					currentDigitString += string(char)
				}

				// we've found the end of the digit (or line), convert the current digit string to a int and then reset all variables
				if !unicode.IsDigit(char) || x == len(row)-1 {
					digit, _ := strconv.Atoi(currentDigitString)
					currentDigitString = ""

					// add the full digit to the global gearPartNumberLists map now we have it
					for coord, _ := range gearsFoundForDigit {
						gearPartNumberLists[coord] = append(gearPartNumberLists[coord], digit)
					}
					gearsFoundForDigit = map[u.Coord]bool{}
				}
			} else {
				// handle outside of digit string
				if unicode.IsDigit(char) {
					_, part, partX, partY := checkNearPartNumber(input, x, y)

					if part == "*" {
						gearsFoundForDigit[u.Coord{X: partX, Y: partY}] = true
					}

					currentDigitString += string(char)
				}
			}
		}
	}

	for _, value := range gearPartNumberLists {
		if len(value) == 2 {
			answer += (value[0] * value[1])
		}
	}
	fmt.Printf("Part Two: %d\n", answer)
}

func checkNearPartNumber(input []string, x int, y int) (bool, string, int, int) {
	for i := y - 1; i <= y+1; i++ {
		if i < 0 || i >= len(input) {
			continue
		}

		for j := x - 1; j <= x+1; j++ {
			if j < 0 || j >= len(input[i]) {
				continue
			}

			if string(input[i][j]) != "." && !unicode.IsDigit(rune(input[i][j])) {
				return true, string(input[i][j]), i, j
			}
		}
	}

	return false, "", 0, 0
}

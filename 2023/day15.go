package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day15.txt"
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

	steps := strings.Split(input[0], ",")
	for _, step := range steps {
		answer += HASH(step)
	}

	fmt.Printf("Part One: %d\n", answer)
}

type Lens struct {
	power int
	label string
}

func partTwo(input []string) {
	// split it, if contains '-' then label, if '=' then label,lensStrength
	// hash the label to get the box number
	// if =, add to the box [label, lensS] if none have label. if they do then replace it
	// if -... go to the box,  box = append(box[:s], box[s+1:]...)
	// then at the end...
	// sum up box_number + 1 * (i+1) + lensPower for every lens in every box
	answer := 0
	boxes := [][]Lens{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, []Lens{})
	}

	steps := strings.Split(input[0], ",")
	for _, step := range steps {
		if strings.Contains(step, "-") {
			label := strings.Split(step, "-")[0]
			boxNumber := HASH(label)
			box := boxes[boxNumber]

			for i, lens := range box {
				if lens.label == label {
					box = append(box[:i], box[i+1:]...)
					boxes[boxNumber] = box
					break
				}
			}
		} else {
			label := strings.Split(step, "=")[0]
			lensStrength, _ := strconv.Atoi(strings.Split(step, "=")[1])
			boxNumber := HASH(label)
			box := boxes[boxNumber]

			foundLens := false
			for i, lens := range box {
				if lens.label == label {
					box[i].power = lensStrength
					foundLens = true
					break
				}
			}
			if !foundLens {
				boxes[boxNumber] = append(box, Lens{label: label, power: lensStrength})
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			answer += (i + 1) * (j + 1) * lens.power
		}
	}

	fmt.Printf("Part Two: %d\n", answer)
}

func HASH(input string) int {
	hash := 0
	for _, char := range input {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day18.txt"
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
	plan := map[u.Coord]string{} //we will store the hex colour code here to indicate its filled in (just "#" for interior), then number of keys is number of dug out squares
	coord := u.NewCoord(0, 0)
	plan[coord] = "#"
	directionMap := map[string]int{
		"U": 0,
		"R": 1,
		"D": 2,
		"L": 3,
	}
	minX, minY, maxX, maxY := 0, 0, 0, 0

	for _, line := range input {
		direction := directionMap[strings.Split(line, " ")[0]]
		length, _ := strconv.Atoi(strings.Split(line, " ")[1])
		colour := strings.Split(line, " ")[2]

		for i := 0; i < length; i++ {
			coord = coord.NewCoordInDirection(direction)
			plan[coord] = colour
		}
		minY = min(minY, coord.Y)
		minX = min(minX, coord.X)
		maxX = max(maxX, coord.X)
		maxY = max(maxY, coord.Y)
	}

	// bit of a hack, taking a point from printed out map thats inside
	queue := []u.Coord{u.NewCoord(1, 0)}
	for len(queue) > 0 {
		nextNode := queue[0]
		queue = queue[1:]

		for _, node := range nextNode.Neighbours(minX, minY, maxX, maxY) {
			if _, ok := plan[node]; !ok {
				queue = append(queue, node)
				plan[node] = "#"
			}
		}
	}

	answer = len(plan)

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// oh dear
	// i dont think part 1 will even work for the example input
	// if we know all the coordinates of the shape formed is there some library or method to calculate the area?
	// answer: apparently yes, 'shoelace' formula, sum up areas of all sub shapes, lets see if it works.
	answer := 0
	coord := u.NewCoord(0, 0)

	for _, line := range input {
		colour := strings.Split(line, " ")[2]
		hexString := colour[2:7]
		distance, _ := strconv.ParseInt(hexString, 16, 64)
		direction, _ := strconv.Atoi(string(colour[7]))

		// not usual distance convention :(
		var newCoord u.Coord
		if direction == 0 {
			newCoord = u.NewCoord(coord.X+int(distance), coord.Y)
		} else if direction == 1 {
			newCoord = u.NewCoord(coord.X, coord.Y-int(distance))
		} else if direction == 2 {
			newCoord = u.NewCoord(coord.X-int(distance), coord.Y)
		} else {
			newCoord = u.NewCoord(coord.X, coord.Y+int(distance))
		}

		answer += ((newCoord.X * coord.Y) - (newCoord.Y * coord.X))
		answer += int(distance) // we need to add the perimeter of the shape to the number of filled in squares as well

		coord = newCoord
	}

	answer = answer / 2
	answer += 1 // add 1 for the initial square filled at (0,0)
	fmt.Printf("Part Two: %d\n", answer)
}

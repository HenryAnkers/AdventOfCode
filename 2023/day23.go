package main

import (
	"fmt"
	"time"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day23.txt"
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
	// just do dfs and find longest path of any neighbours
	answer := 0
	start := u.NewCoord(1, 0)
	answer = longestPathLength(input, start, map[u.Coord]bool{start: true}, true) - 1

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// part 1 is only 70ms so I can't imagine part2 is that hard to bruteforce but...
	// Part One took: 70.0944ms, Part Two took: 1h19m39.5997935s
	// I think we can cut down significantly on the number of paths by only considering junction nodes and the length between them
	// maybe I implement later
	answer := 0
	start := u.NewCoord(1, 0)
	answer = longestPathLength(input, start, map[u.Coord]bool{start: true}, false) - 1
	fmt.Printf("Part Two: %d\n", answer)
}

func longestPathLength(input []string, position u.Coord, visited map[u.Coord]bool, part1 bool) int {
	visited[position] = true

	if position.Y == len(input)-1 && position.X == len(input[0])-2 {
		return 1
	}

	currentChar := input[position.Y][position.X]
	potentialNeighbours := position.Neighbours(0, 0, len(input[0])-1, len(input)-1)
	if part1 {
		switch currentChar {
		case '^':
			potentialNeighbours = []u.Coord{position.North()}
		case '>':
			potentialNeighbours = []u.Coord{position.East()}
		case 'v':
			potentialNeighbours = []u.Coord{position.South()}
		case '<':
			potentialNeighbours = []u.Coord{position.West()}
		}
	}

	maxAdditionalPathLengthFound := 0
	for _, neighbourNode := range potentialNeighbours {
		character := input[neighbourNode.Y][neighbourNode.X]
		if visited[neighbourNode] || character == '#' {
			continue
		}

		if neighbourNode.Y == len(input)-1 && neighbourNode.X == len(input[0])-2 {
			return 2
		}

		visited[neighbourNode] = true
		potentialLongest := longestPathLength(input, neighbourNode, visited, part1)
		visited[neighbourNode] = false
		maxAdditionalPathLengthFound = max(maxAdditionalPathLengthFound, potentialLongest)
	}
	if maxAdditionalPathLengthFound == 0 {
		// We have no more nodes to visit but we haven't hit the end
		return 0
	}

	return 1 + maxAdditionalPathLengthFound
}

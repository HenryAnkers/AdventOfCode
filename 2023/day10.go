package main

import (
	"fmt"
	"math"
	"time"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day10.txt"
	input, err := u.ReadLines(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start := time.Now()
	// how do we detect a loop?
	// - do a DFS
	// - check at every node if from this node we can access one already visited, if so we've found the beginning of a loop at the previously visited nodes
	// - if we do have a loop, then we need to find the distance. we can do this by keeping track of the parents of each node and iterating back until we find the previously visited node again
	// - if not, add the node to visited (visited[node] = parentNode), add it to list of nodes to process, continue
	// can we just start DFS at (0,0) and find it? no we need to find S first...

	answer := 0
	parentNodes := map[u.Coord]u.Coord{}
	distances := map[u.Coord]int{} // lazy
	queue := []u.Coord{}
	loopCoords := []u.Coord{}
	startCoord := u.Coord{}

findStartPosition:
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				startCoord.X = x
				startCoord.Y = y
				queue = append(queue, startCoord)
				break findStartPosition
			}
		}
	}

	parentNodes[startCoord] = u.Coord{X: -1, Y: -1}
	distances[startCoord] = 0

findLoop:
	// do DFS until we find the loop or run out of nodes...
	for len(queue) > 0 {
		coord := queue[0]
		queue = queue[1:]
		x, y := coord.X, coord.Y
		currentChar := input[y][x]
		parentX, parentY := parentNodes[coord].X, parentNodes[coord].Y

		// | is a vertical pipe connecting north and south.
		// - is a horizontal pipe connecting east and west.
		// L is a 90-degree bend connecting north and east.
		// J is a 90-degree bend connecting north and west.
		// 7 is a 90-degree bend connecting south and west.
		// F is a 90-degree bend connecting south and east.
		// . is ground; there is no pipe in this tile.
		// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

		coordsToCheck := []u.Coord{}
		if currentChar == '|' {
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.North(), coord.South()}...)
		} else if currentChar == '-' {
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.East(), coord.West()}...)
		} else if currentChar == 'L' {
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.North(), coord.East()}...)
		} else if currentChar == 'J' {
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.North(), coord.West()}...)
		} else if currentChar == '7' {
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.South(), coord.West()}...)
		} else if currentChar == 'F' {
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.South(), coord.East()}...)
		} else if currentChar == 'S' {
			// as we dont know the shape of the pipe we don't know which are valid moves... is there a better way than checking manually?
			coordsToCheck = append(coordsToCheck, []u.Coord{coord.South(), coord.North()}...)
		}

		for _, newCoord := range coordsToCheck {
			newX, newY := newCoord.X, newCoord.Y

			if !(newY < 0 || newY >= len(input) || newX < 0 || newX >= len(input[newY]) || input[newY][newX] == '.' || (newX == parentX && newY == parentY)) {
				// we know the node isnt this nodes parent, so if its been visited we've found a loop
				if _, ok := parentNodes[newCoord]; ok {
					// add the path from newCoord back to startingCoord in reverse, so we have S -> ... -> newCoord
					pointer := newCoord
					for pointer.X != startCoord.X || pointer.Y != startCoord.Y {
						loopCoords = append([]u.Coord{pointer}, loopCoords...)
						pointer = parentNodes[pointer]
					}
					loopCoords = append([]u.Coord{startCoord}, loopCoords...)

					// add the path from current coord back to startingCoord to the end of the list, so we have S -> ... -> newCoord -> coord -> ... -> (node before S)
					pointer = coord
					for pointer.X != startCoord.X || pointer.Y != startCoord.Y {
						loopCoords = append(loopCoords, pointer)
						pointer = parentNodes[pointer]
					}

					answer = int(math.Ceil(float64(len(loopCoords)) / 2))
					fmt.Println("Part One:", answer)
					break findLoop
				} else {
					parentNodes[newCoord] = coord
					distances[newCoord] = distances[coord] + 1
					queue = append([]u.Coord{newCoord}, queue...)
				}
			}
		}
	}
	duration1 := time.Since(start)

	// part two
	// we now need to find the boundary the loop encompasses
	// iterate through the input, find all "." characters (no... all coords not inside the loop we found	)
	// do a BFS every time we find one, determine if the island is inside or outside the loop (if any coordinates in the island == loop boundary, then its outside the loop)
	// add to total if so
	// BUT WE CAN SQUEEZE BETWEEN PIPES FFS

	// ..........
	// .F------7.
	// .|F----7|.
	// .||OOOO||.
	// .||OOOO||.
	// .|L-7F-J|.
	// .|II||II|.
	// .L--JL--J.
	// ..........
	// can see why OOOOO block isn't in the loop despite being within the coordinates...
	// scan left to right, every time we see a break in the loop (i.e. the loop stops moving horizontally), toggle if we're in the boundary or not
	// break in the loop = we find a vertical pipe '|', or we find a part where we have L---7 or F---J ... basically the line opens in a different direction to it closes (north-south)
	// four cases: F---7 no change, L---J no change, L---7 change, F---J change
	// so basically change whenever you see L or J
	// ...or the starting character, which must be J in our case

	loopCoordsMap := map[u.Coord]bool{}
	for _, v := range loopCoords {
		loopCoordsMap[v] = true
	}

	enclosedByLoop := 0
	for y, line := range input {
		isInLoop := false

		for x, char := range line {
			coord := u.Coord{X: x, Y: y}
			if loopCoordsMap[coord] && (char == '|' || char == 'L' || char == 'J' || char == 'S') {
				isInLoop = !isInLoop
			}

			if isInLoop && !loopCoordsMap[coord] {
				enclosedByLoop++
			}
		}
	}

	fmt.Println("Part Two:", enclosedByLoop)
	duration2 := time.Since(start)
	fmt.Printf("Part One took: %v, Part Two took: %v\n", duration1, duration2)
}

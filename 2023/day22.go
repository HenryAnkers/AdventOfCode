package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day22.txt"
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
	// this is like the tetris last year but worse?
	// for each shape we basically have to go down until we hit a colision based on the x,y shape (the shapes share at least one point)
	// we will order by z coordinate and process the lowest ones first
	// how do we determine which can be removed? its all shapes that are resting on exactly one other shape we need to find (specifically what the other shape is). although we only care about the count so it doesnt matter
	// at every step find every collision, if theres none move it down and continue, else save how many different shapes (we'll save them as diff strings in map) collide with the object (these objects are further down so will have been processed beforehand)
	answer := 0
	allShapes, allShapeCoords := parseShapesFromInput(input)
	_, supportedBy := processGravity(allShapes, allShapeCoords)

	nodesWeCantDisintegrate := map[int]bool{}
	for _, v := range supportedBy {
		if len(v) == 1 { // this node is only supported by one other node, so we cant remove that node
			nodesWeCantDisintegrate[v[0]] = true
		}
	}
	answer = len(allShapes) - len(nodesWeCantDisintegrate)

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// when we remove a brick that is the only support for another brick, that second brick falls
	// if the first brick we remove removes 2, and those 2 are the only supports... whats a nice way of modelling this?
	// we might need to map it both way, so we have
	// supporting[a] = [b,c]
	// supportedBy[b],[c] = a
	// then we queue a for removal, for every item B in supporting[a], remove a from supportedBy[B]. If this leaves the list empty then add to the queue
	answer := 0
	allShapes, allShapeCoords := parseShapesFromInput(input)
	supporting, supportedBy := processGravity(allShapes, allShapeCoords)

	processedSingleNodes := map[int]bool{}
	for _, v := range supportedBy {
		// this node is only supported by one other node, lets try removing it
		if len(v) == 1 {
			if processedSingleNodes[v[0]] {
				continue
			}
			processedSingleNodes[v[0]] = true
			currentRemoved := map[int]bool{}
			currentRemoved[v[0]] = true
			queue := []int{v[0]}

			for len(queue) > 0 {
				newQueue := []int{}

				nodesThatWereSupported := []int{}
				// find all nodes that nodes in this queue were supporting
				for _, removedNode := range queue {
					nodesThatWereSupported = append(nodesThatWereSupported, supporting[removedNode]...)
				}

				// for each node that was being supported by at least one removed node, we check to see if ALL the supporting node have been removed, and add them to the new queue if so
				for _, node := range nodesThatWereSupported {
					allSupportsRemoved := true
					for _, supportingNode := range supportedBy[node] {
						if !currentRemoved[supportingNode] {
							allSupportsRemoved = false
							break
						}
					}
					if allSupportsRemoved {
						currentRemoved[node] = true
						newQueue = append(newQueue, node)
					}
				}
				queue = newQueue
			}

			answer += len(currentRemoved) - 1
		}
	}

	fmt.Printf("Part Two: %d\n", answer)
}

func parseShapesFromInput(input []string) ([][]u.Coord, map[u.Coord]int) {
	shapes := [][]u.Coord{}

	for _, line := range input {
		firstCoordList := strings.Split(strings.Split(line, "~")[0], ",")
		secondCoordList := strings.Split(strings.Split(line, "~")[1], ",")
		allCoords := []int{}
		for _, c := range firstCoordList {
			converted, _ := strconv.Atoi(c)
			allCoords = append(allCoords, converted)
		}
		for _, c := range secondCoordList {
			converted, _ := strconv.Atoi(c)
			allCoords = append(allCoords, converted)
		}

		newShape := []u.Coord{}
		newShape = append(newShape, u.NewCoord3D(allCoords[0], allCoords[1], allCoords[2]))
		newShape = append(newShape, u.NewCoord3D(allCoords[3], allCoords[4], allCoords[5]))
		shapes = append(shapes, newShape)
	}

	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i][0].Z < shapes[j][0].Z
	})

	allShapeCoords := map[u.Coord]int{} //int maps to shape index above +1
	for i, shape := range shapes {
		lowX, highX := shape[0].X, shape[1].X
		lowY, highY := shape[0].Y, shape[1].Y
		lowZ, highZ := shape[0].Z, shape[1].Z

		for x := lowX; x <= highX; x++ {
			for y := lowY; y <= highY; y++ {
				for z := lowZ; z <= highZ; z++ {
					allShapeCoords[u.NewCoord3D(x, y, z)] = i + 1
				}
			}
		}
	}

	return shapes, allShapeCoords
}

func processGravity(allShapes [][]u.Coord, allShapeCoords map[u.Coord]int) (map[int][]int, map[int][]int) {
	supporting := map[int][]int{}  //the nodes that map[int] is supporting
	supportedBy := map[int][]int{} // the nodes that map[int] is supported by
	for i, _ := range allShapes {
		supportedBy[i+1] = []int{}
		supporting[i+1] = []int{}
	}

processShapes:
	for i, shape := range allShapes {
		for true {
			currentShapeIndex := i + 1
			lowZ := min(shape[0].Z, shape[1].Z)
			if lowZ <= 1 {
				continue processShapes
			}

			// check all coords with z-1. if we find other shapes at this level, see how many there are
			lowX, highX := shape[0].X, shape[1].X
			lowY, highY := shape[0].Y, shape[1].Y
			lowZ, highZ := shape[0].Z, shape[1].Z

			uniqueCollisions := map[int]int{}
			for x := lowX; x <= highX; x++ {
				for y := lowY; y <= highY; y++ {
					for z := lowZ; z <= highZ; z++ {
						if shapeInt, ok := allShapeCoords[u.NewCoord3D(x, y, z-1)]; ok {
							if shapeInt != currentShapeIndex && shapeInt != 0 {
								uniqueCollisions[shapeInt] += 1
							}
						}
					}
				}
			}
			if len(uniqueCollisions) > 0 {
				for key, _ := range uniqueCollisions {
					supportedBy[currentShapeIndex] = append(supportedBy[currentShapeIndex], key)
					supporting[key] = append(supporting[key], currentShapeIndex)
				}

				continue processShapes
			}

			// we havent found a collision, lets move it down
			for x := lowX; x <= highX; x++ {
				for y := lowY; y <= highY; y++ {
					for z := lowZ; z <= highZ; z++ {
						allShapeCoords[u.NewCoord3D(x, y, z-1)] = currentShapeIndex
						delete(allShapeCoords, u.NewCoord3D(x, y, z))
					}
				}
			}
			shape[0].Z -= 1
			shape[1].Z -= 1
		}
	}

	return supporting, supportedBy
}

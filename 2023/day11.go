package main

import (
	"fmt"
	"time"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day11.txt"
	input, err := u.ReadLines(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start := time.Now()
	answer(input)
	duration1 := time.Since(start)
	fmt.Printf("Took: %v", duration1)
}

func answer(input []string) {
	answer := 0
	answertwo := 0
	// iterate over the list, marking coordinates of stars
	// IF the row or column is empty (deal with column later), then we mark expandedRow/expandedColumn[x/y] as true
	// we then calculate the distances and see how many expandedrows/columns we are forced to go through
	// answer for that pair is distance+numberofexpanded

	columnHasStar := map[int]bool{}
	rowHasStar := map[int]bool{}
	starMap := []u.Coord{}

	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				columnHasStar[x] = true
				rowHasStar[y] = true
				starMap = append(starMap, u.NewCoord(x, y))
			}
		}
	}

	for i := 0; i < len(starMap); i++ {
		for j := i + 1; j < len(starMap); j++ {
			starI, starJ := starMap[i], starMap[j]
			maxX, minX := max(starI.X, starJ.X), min(starI.X, starJ.X)
			maxY, minY := max(starI.Y, starJ.Y), min(starI.Y, starJ.Y)
			rawDistance := (maxX - minX) + (maxY - minY)

			expandedDistance := 0
			for d := minX; d <= maxX; d++ {
				if !columnHasStar[d] {
					expandedDistance++
				}
			}
			for d := minY; d <= maxY; d++ {
				if !rowHasStar[d] {
					expandedDistance++
				}
			}

			answer += rawDistance + expandedDistance
			answertwo += rawDistance + (expandedDistance * (1000000 - 1))
		}
	}

	fmt.Println("Part One:", answer)
	fmt.Println("Part Two:", answertwo)

}

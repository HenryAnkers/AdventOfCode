package main

import (
	"fmt"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day13.txt"
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
	// map each row and column seperately
	// iterate through row index, see if any is a mirror, if not repeat with column
	// for every row/column index, check if (i, i+1) is symmetrical whilst this is a valid set of coordinates
	// if all true until we go off one end, then this is a mirror
	answer := 0

	grids := [][]string{}
	currentGrid := []string{}
	for _, line := range input {
		if line == "" {
			grids = append(grids, currentGrid)
			currentGrid = []string{}
		} else {
			currentGrid = append(currentGrid, line)
		}
	}
	grids = append(grids, currentGrid)

processGrid:
	for _, grid := range grids {
		rows := []string{}
		columns := []string{}

		for _, line := range grid {
			rows = append(rows, line)
		}
		for x, _ := range grid[0] {
			currentStr := ""
			for y := 0; y < len(grid); y++ {
				currentStr += string(grid[y][x])
			}
			columns = append(columns, currentStr)
		}

		for y := 0; y < len(rows)-1; y++ {
			isMirror := true
			for j := 0; (y+j+1 < len(rows)) && (y >= j); j++ {
				if rows[y+j+1] != rows[y-j] {
					isMirror = false
					break
				}
			}

			if isMirror {
				answer += 100 * (y + 1)
				continue processGrid
			}
		}

		for x := 0; x < len(columns)-1; x++ {
			isMirror := true
			for j := 0; (x+j+1 < len(columns)) && (x >= j); j++ {
				if columns[x+j+1] != columns[x-j] {
					isMirror = false
					break
				}
			}

			if isMirror {
				answer += x + 1
				continue processGrid
			}
		}
	}

	fmt.Println("Part One:", answer)
}

func partTwo(input []string) {
	// ffs...
	// we need to do the same check but instead look for a single pair which doesnt match but has an edit distance of exactly 1
	// i guess the total edit distance of all pairs should therefore be one, so we can just track overall diff
	answer := 0

	grids := [][]string{}
	currentGrid := []string{}
	for _, line := range input {
		if line == "" {
			grids = append(grids, currentGrid)
			currentGrid = []string{}
		} else {
			currentGrid = append(currentGrid, line)
		}
	}
	grids = append(grids, currentGrid)

processGrid:
	for _, grid := range grids {
		rows := []string{}
		columns := []string{}

		for _, line := range grid {
			rows = append(rows, line)
		}
		for x, _ := range grid[0] {
			currentStr := ""
			for y := 0; y < len(grid); y++ {
				currentStr += string(grid[y][x])
			}
			columns = append(columns, currentStr)
		}

		for y := 0; y < len(rows)-1; y++ {
			totalDiff := 0
			for j := 0; (y+j+1 < len(rows)) && (y >= j); j++ {
				totalDiff += getEditDistance(rows[y+j+1], rows[y-j])
				if totalDiff > 1 {
					break
				}
			}

			if totalDiff == 1 {
				answer += 100 * (y + 1)
				continue processGrid
			}
		}

		for x := 0; x < len(columns)-1; x++ {
			totalDiff := 0
			for j := 0; (x+j+1 < len(columns)) && (x >= j); j++ {
				totalDiff += getEditDistance(columns[x+j+1], columns[x-j])
				if totalDiff > 1 {
					break
				}
			}

			if totalDiff == 1 {
				answer += x + 1
				continue processGrid
			}
		}
	}

	fmt.Println("Part Two:", answer)
}

func getEditDistance(a string, b string) int {
	ans := 0
	for x, _ := range a {
		if a[x] != b[x] {
			ans++
		}
	}
	return ans
}

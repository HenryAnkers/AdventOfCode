package main

import (
	"fmt"
	"strings"
	"time"

	"adventofcode/utils"
	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day16.txt"
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
	answer := illuminate(input, []Beam{Beam{
		coord:     u.NewCoord(-1, 0),
		direction: 1,
	}})

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	answer := 0

	for y := 0; y < len(input); y++ {
		answer = max(answer,
			illuminate(input, []Beam{Beam{
				coord:     u.NewCoord(-1, y),
				direction: 1,
			}}))
	}
	for x := 0; x < len(input[0]); x++ {
		answer = max(answer,
			illuminate(input, []Beam{Beam{
				coord:     u.NewCoord(x, -1),
				direction: 2,
			}}))
	}

	fmt.Printf("Part Two: %d\n", answer)
}

type Beam struct {
	coord u.Coord
	// 0123 upleftdownright
	direction int
}

func getCacheKey(beam Beam) string {
	return strings.Join([]string{string(beam.coord.X), string(beam.coord.Y), string(beam.direction)}, " - ")
}

func illuminate(input []string, startingBeams []Beam) int {
	energized := map[u.Coord]bool{}
	memo := map[string]bool{}
	beams := startingBeams

	for len(beams) > 0 {
		newBeams := []Beam{}
		for _, beam := range beams {
			key := getCacheKey(beam)
			if memo[key] {
				continue
			}
			energized[beam.coord] = true

			result := []Beam{}
			if beam.direction == 0 {
				result = processMove(input, beam.coord.North(), 0)
			} else if beam.direction == 1 {
				result = processMove(input, beam.coord.East(), 1)
			} else if beam.direction == 2 {
				result = processMove(input, beam.coord.South(), 2)
			} else {
				result = processMove(input, beam.coord.West(), 3)
			}

			newBeams = append(newBeams, result...)
			memo[key] = true
		}

		beams = newBeams
	}

	return (len(energized) - 1) // -1 as starting coordinate is never valid
}

func processMove(input []string, newCoord u.Coord, direction int) []Beam {
	x, y := newCoord.X, newCoord.Y
	if y < 0 || x < 0 || y >= len(input) || x >= len(input[y]) {
		return []Beam{}
	}

	symbol := input[y][x]
	if symbol == '/' {
		// 0 -> 1, 1 -> 0, 3 -> 2, 2 -> 3
		newDirection := 0
		if direction == 0 || direction == 2 {
			newDirection = direction + 1
		} else {
			newDirection = direction - 1
		}

		return []Beam{Beam{
			coord:     newCoord,
			direction: newDirection,
		}}
	} else if symbol == '\\' {
		// 0 -> 3, 3 -> 0, 1 -> 2, 2 -> 1
		newDirection := 0
		if direction == 0 {
			newDirection = 3
		} else if direction == 3 {
			newDirection = 0
		} else if direction == 1 {
			newDirection = 2
		} else {
			newDirection = 1
		}

		return []Beam{Beam{
			coord:     newCoord,
			direction: newDirection,
		}}
	} else if symbol == '-' {
		if direction == 0 || direction == 2 {
			newBeams := []Beam{}
			newBeams = append(newBeams, Beam{
				coord:     u.NewCoord(x, y),
				direction: 1,
			})
			newBeams = append(newBeams, Beam{
				coord:     u.NewCoord(x, y),
				direction: 3,
			})
			return newBeams
		}
	} else if symbol == '|' {
		if direction == 1 || direction == 3 {
			newBeams := []Beam{}
			newBeams = append(newBeams, Beam{
				coord:     u.NewCoord(x, y),
				direction: 2,
			})
			newBeams = append(newBeams, Beam{
				coord:     u.NewCoord(x, y),
				direction: 0,
			})
			return newBeams
		}
	}

	return []Beam{Beam{
		coord:     newCoord,
		direction: direction,
	}}
}

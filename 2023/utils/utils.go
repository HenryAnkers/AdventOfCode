package utils

import (
	"bufio"
	"math"
	"os"
)

type Coord struct {
	X int
	Y int
}

func (c Coord) South() Coord {
	return Coord{X: c.X, Y: c.Y + 1}
}

func (c Coord) North() Coord {
	return Coord{X: c.X, Y: c.Y - 1}
}

func (c Coord) East() Coord {
	return Coord{X: c.X + 1, Y: c.Y}
}

func (c Coord) West() Coord {
	return Coord{X: c.X - 1, Y: c.Y}
}

func (c Coord) Neighbours(minX, minY, maxX, maxY int) []Coord {
	allNeighbours := []Coord{c.North(), c.East(), c.South(), c.West()}
	validNeighbours := []Coord{}

	for _, n := range allNeighbours {
		if !(n.X < minX || n.X > maxX || n.Y < minY || n.Y > maxY) {
			validNeighbours = append(validNeighbours, n)
		}
	}

	return validNeighbours
}

func NewCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

// 0123 northeastsouthwest is convention used for what direction we're pointing in
func (c Coord) NewCoordInDirection(direction int) Coord {
	if direction == 0 {
		return c.North()
	}
	if direction == 1 {
		return c.East()
	}
	if direction == 2 {
		return c.South()
	}
	return c.West()
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return int(math.Abs(float64(a*b))) / GCD(a, b)
}

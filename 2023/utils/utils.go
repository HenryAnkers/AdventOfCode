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

func NewCoord(x int, y int) Coord {
	return Coord{X: x, Y: y}
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

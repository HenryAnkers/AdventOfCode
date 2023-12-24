package main

import (
	"fmt"
	"strings"
	"time"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day24.txt"
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
	// > ignore the Z axis
	// not looking forward to part 2...
	// for part 1 we can model these equations in the format y = mx + c, we have (y_a = m_a*x + c_a), (y_b = m_b*x + c_b), intersecting at (x_i,y_i), then:
	// m_a(x_i) + c_a = m_b(x_i) + c_b = y_i; m_a(x_a) - m_b(x_i) = c_b - c_a; x_i = (c_b-c_a)/(m_a-m_b) obviously if m_a = m_b then they wont intersect ever unless same equation, otherwise this gives us intersection point
	answer := 0
	allEquations := parseInput(input)
	start, end := 200000000000000.0, 400000000000000.0

	for i := 0; i < len(allEquations); i++ {
		equationA := allEquations[i]
		x_a, y_a := float64(equationA[0][0]), float64(equationA[0][1])
		dX, dY := float64(equationA[1][0]), float64(equationA[1][1])
		m_a := float64(dY) / float64(dX)
		c_a := y_a - m_a*x_a

		for j := i + 1; j < len(allEquations); j++ {
			equationB := allEquations[j]
			x_b, y_b := float64(equationB[0][0]), float64(equationB[0][1])
			dXb, dYb := float64(equationB[1][0]), float64(equationB[1][1])
			m_b := float64(dYb) / float64(dXb)
			c_b := y_b - m_b*x_b

			if m_a == m_b {
				continue
			}

			intersectionX := (c_b - c_a) / (m_a - m_b)
			intersectionY := m_a*intersectionX + c_a
			intersectionInFutureForA := dX < 0 && intersectionX < x_a || dX > 0 && intersectionX > x_a
			intersectionInFutureForB := dXb < 0 && intersectionX < x_b || dXb > 0 && intersectionX > x_b

			if intersectionX >= start && intersectionY >= start && intersectionX <= end && intersectionY <= end && intersectionInFutureForA && intersectionInFutureForB {
				answer += 1
			}
		}
	}

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// >Now including the Z axis - who could have guessed
	// we need to throw a rock which intersects every single other line.... oh dear
	// we can model this as a set of three equations (x = x_c + T*dX, y = y_c + T*dY, z_t = z_c + T*dZ)
	// let our rock have properties x = x_r + T*dXr, y = y_r + T+dYr, z = z_r + T*dZr
	// then at the intersection x_i. y_i, z_i at T0;  x_r + T0*dXr = x_c + T0*dX = x_i;   (x_r + T0*(dXr-dX)) = x_c (same for Y and Z)
	// we are left with a huge list of equations in x_r and dXr for various values of T (which we don't know)...  there must be a better way? How can we get something we can plug into Z3 at worst case?

	//  (x_0 - x_i) = T_n(dXr - dXi) for every rock...
	//  X  = Xi + T_n(dX - dXi) -> there is a position that all these equations pass through with modified velocity for every position Xi
	// Can we somehow fix velocity and iterate over these possiiblites (assuming velocity <1k cos all the rocks are).
	// Now we can calculate the RHS of the equation and get a system A*Tn + B = X, if this is solvable with all Tn > 0 then we have found an answer. If its not solveable (how do we tell???) then we move on to the next velocity
	// if some pair of bailstones intersect at some differemt point, then they wont intersect at X, but all the others will.. implies this is true if the intersection point for all rocks is the same => intersection point for all rocks always exist and always same coord
	// => find the intersection point (x,y,z) for every coord, if we find more than one intersection point for two different pairs of coords then we break and move onto next velocity
	// how do we convert part 1 code into finding intersection point x,y,z? if the x,y intersection time is the same as the y,z intersection time?
	// idk and I dont have time to find out. maybe will come back to his later

}

func parseInput(input []string) [][][]int {
	output := [][][]int{}
	for _, line := range input {
		initialCoords := u.ConvertStrIntArray(strings.Split(strings.Split(line, " @ ")[0], ", "))
		deltas := u.ConvertStrIntArray(strings.Split(strings.Split(line, " @ ")[1], ", "))

		output = append(output, [][]int{initialCoords, deltas})
	}

	return output
}

// allEquations := parseInput(input)
// allGrads := map[float64]map[float64]int{}

// for i := 0; i < len(allEquations); i++ {
// 	equationA := allEquations[i]
// 	dX, dY, dZ := float64(equationA[1][0]), float64(equationA[1][1]), float64(equationA[1][2])
// 	m_a_x := float64(dY) / float64(dX)
// 	m_a_z := float64(dY) / float64(dZ)

// 	if m, ok := allGrads[m_a_x]; ok {
// 		if v, ok := m[m_a_z]; ok {
// 			fmt.Println(equationA)
// 			fmt.Println(allEquations[v])
// 			fmt.Println()
// 		}
// 	} else {
// 		allGrads[m_a_x] = map[float64]int{}
// 	}
// 	allGrads[m_a_x][m_a_z] = i
// }

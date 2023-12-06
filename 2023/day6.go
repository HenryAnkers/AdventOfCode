package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Iterative solution
	times := []int{49, 97, 94, 94}
	distances := []int{263, 1532, 1378, 1851}
	start := time.Now()
	solution(times, distances)
	duration1 := time.Since(start)

	times = []int{49979494}
	distances = []int{263153213781851}
	start = time.Now()
	solution(times, distances)
	duration2 := time.Since(start)
	fmt.Printf("Normal: Part One took: %v, Part Two took: %v\n", duration1, duration2)

	// Solving quadratic equation
	times = []int{49, 97, 94, 94}
	distances = []int{263, 1532, 1378, 1851}
	start = time.Now()
	quadratic(times, distances)
	duration3 := time.Since(start)

	times = []int{49979494}
	distances = []int{263153213781851}
	start = time.Now()
	quadratic(times, distances)
	duration4 := time.Since(start)
	fmt.Printf("Quadratic: Part One took: %v, Part Two took: %v\n", duration3, duration4)
}

func solution(times []int, distances []int) {
	// let total time be T, time held = k_s
	// speed is t * k_s, so distance is k_s * (T-k_s)

	answer := 1

	for i, time := range times {
		previousMaxDistance := distances[i]
		localResult := 0

		for timeHeld := 0; timeHeld <= time; timeHeld++ {
			result := timeHeld * (time - timeHeld)
			if result > previousMaxDistance {
				localResult += 1
			}
		}

		if localResult > 0 {
			answer *= localResult
		}
	}

	fmt.Printf("Answer: %d\n", answer)
}

func quadratic(times []int, distances []int) {
	// from above, we want to know how many values for k_s the inequality k_s * (T-k_s) > Distance holds
	// k_s * (T-k_s) > distance
	// k_s * (T-k_s) - distance > 0
	// -k_s^2 + T*k_s - distance > 0
	// k_s^2 - T*k_s + distance < 0
	// find what values of k_s this is valid for... quadratic equation, find two zero points
	// it has to be a range (x,y) where 0 < x < y < T just by the nature of the question, if not we wouldnt be able to get an answer
	// how do we deal with roots that are decimals? round up lower bound, round down upper bound, but then it's inclusive on the lower bound so add one

	answer := 1

	for i, time := range times {
		distance := distances[i]
		timef := float64(time)
		distancef := float64(distance)

		// quadratic equation (for ax^2 + bx + c): X = (-b +/- sqrt(b^2 - 4ac))/2a
		root1 := (timef + math.Sqrt((timef*timef)-(4*distancef))) / 2
		root2 := (timef - math.Sqrt((timef*timef)-(4*distancef))) / 2

		localAnswer := math.Floor(max(root1, root2)) - math.Ceil(min(root1, root2)) + 1
		answer *= int(localAnswer)
	}

	fmt.Printf("Answer: %d\n", answer)
}

package main

import (
	"fmt"
	"math"
	"time"

	u "adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day21.txt"
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
	// just do BFS and return the end queue?
	// i imagine part 2 is we take a billion steps and we need to look for cycles again?
	answer := 0
	startingCoord := getStartingCoord(input)
	queue := []u.Coord{startingCoord}

	for steps := 0; steps < 64 && len(queue) > 0; steps++ {
		newQueueMap := map[u.Coord]bool{}

		for _, coordToProcess := range queue {
			potentialCoords := coordToProcess.Neighbours(0, 0, len(input[0])-1, len(input)-1)
			for _, c := range potentialCoords {
				if input[c.Y][c.X] != '#' {
					newQueueMap[c] = true
				}
			}
		}

		newQueue := []u.Coord{}
		for k, _ := range newQueueMap {
			newQueue = append(newQueue, k)
		}
		queue = newQueue
	}
	answer = len(queue)

	fmt.Printf("Part One: %d\n", answer)
}

type Step struct {
	absoluteCoord u.Coord
	relativeCoord u.Coord
}

func partTwo(input []string) {
	// this looks even worse than yesterday, this might be it for me this year :(
	// we can deal with the map repeating by just teleporting them to the other side, we need to store their actual & relative coordinates
	// even if we have this, how do we deal with the huge number of steps?
	// we cant do a standard BFS and find all reachable nodes, because I think the queue would become too large (could double every iteration or even 4x?)
	// 6 steps = 15, 50 = 1594, 500 = 167004, 5000 = 16733044 seems to suggest theres probably some sort of pattern/cycle here? increase by 10 seems to increase the number of steps by approx 100 (=> x^2 but does this hold for the real input, dont think so)
	// at one step x_i with K squares, the next step you end with x_i^(1-4). None of this really helps us find an answer though...
	// one thing: part 1 will eventually get to a pattern of (7748, 7757, 7748...) regardless of the input length.  IF we only consider the relative coordinates to the actual input map, we should still reach some equilibrium (i think) in part 2
	// once we have that... if we enter a step and we've been in this relative coordinate before BUT THE ACTUAL COORDINATE IS DIFFERENT  then we already know all the places it can reach in X steps where X is the difference between the step where we first saw it and now
	// can we just save the relative board state and keep track of the steps every time we see the same state instead? I guess one worry is having 1,2,4... pieces on the same relative coordinate :(
	// as long as they all increased proportionally it's fine I guess?

	// instead can we just track the rate of change d(number of positions)/d(number of steps) between 1000s of steps until it becomes a stable value? then just work out how many more we make till the end somehow...
	// its not neccesarily stable. we will have the same pattern but more and more steps on each relative coord, worst case is its some equation x^4
	// but if the equation is ^4 then (26501365^4) > int64 so i sort of hope this isnt the case even for real input, lets see
	// lets just code and look at some values up to 1k

	answer := 0
	startingCoord := getStartingCoord(input)
	queue := []u.Coord{startingCoord}
	stepValues := map[int]int{}

	for steps := 0; steps < 328; steps++ {
		newAbsoluteQueueMap := map[u.Coord]bool{}
		for _, stepToProcess := range queue {
			potentialAbsoluteCoords := stepToProcess.Neighbours(-math.MaxInt64, -math.MaxInt64, math.MaxInt64, math.MaxInt64)

			for _, c := range potentialAbsoluteCoords {
				newRelativeY := (c.Y) % len(input)
				if newRelativeY < 0 {
					newRelativeY += len(input)
				}

				newRelativeX := (c.X) % len(input[0])
				if newRelativeX < 0 {
					newRelativeX += len(input[0])
				}

				if input[newRelativeY][newRelativeX] != '#' {
					newAbsoluteQueueMap[c] = true
				}
			}
		}

		newQueue := []u.Coord{}
		for k, _ := range newAbsoluteQueueMap {
			newQueue = append(newQueue, k)
		}
		queue = newQueue
		stepValues[steps+1] = len(queue)
	}

	// fmt.Println(stepValues)
	// after the 50th step, lets see the delta between each step divided by the stepvalue
	// seems to vary between positive and negative but trends upwards...
	// by just printing out a 0 if the delta is negative, 1 if its positive, we get (for the example input)
	// 0011011101010101110011001010100111010101001110101010011101010100111010101001110101010011101010100111010101001110101010011101010100111010101001110101010011101010100111010101001110101010011101010100111010101001110101010011101010100111010101001110101010011101010100111010101001110101010011101010100111
	// which we can see is the string 10011101010 repeated thanks to visual studio code, 11 characters long, so lets look at the delta after 11 steps

	// -- Messing around with raw delta values between steps, printing delta string --
	// deltas := map[int][]float64{}
	// allDeltaPositiveNegString := ""
	// for k := 1; k < len(stepValues); k++ {
	// 	v := stepValues[k]
	// 	currentVal := v
	// 	prevVal := stepValues[k-1]
	// 	delta := float64(currentVal - prevVal)
	// 	deltas[k] = []float64{delta}
	// 	prevDelta, ok := deltas[k-1]
	// 	if ok {
	// 		temp := 0
	// 		if delta-prevDelta[0] > 0 {
	// 			temp = 1
	// 		}
	// 		allDeltaPositiveNegString += strconv.Itoa(temp)
	// 		deltas[k] = append(deltas[k], float64(temp))
	// 	}
	// }
	// fmt.Println(deltas)
	// fmt.Println(allDeltaPositiveNegString)

	// now trying with 11 gap between steps...
	// WE GET A CONSISTENT DIFFERENCE OF 162 BETWEEN THE DELTA BETWEEN 11 STEPS AFTER A WHILE!
	// what does this mean...?
	// given some arbitary start point like 64 = 2665, delta = 860, we should expect to see 75 = 2665+(860 + 162*1)=3687, 86 = 3687+(860 + 162*2)=4871, which we do
	// we do not see this for our real input..
	// ok so this worked for the example input but not the real one. the difference has to be due to the grid size. is 11 relevant to the input at all? its the literal width of the grid, so lets try new the new width (131)
	// WE GET 31010!

	// -- Checking Delta after k=11 steps for example, 131 steps for real input --
	// deltas := map[int][]float64{}
	// for k := 131; k < len(stepValues); k++ {
	// 	v := stepValues[k]
	// 	currentVal := v
	// 	prevVal := stepValues[k-131]
	// 	delta := float64(currentVal - prevVal)
	// 	deltas[k] = []float64{delta}
	// 	prevDelta, ok := deltas[k-131]
	// 	if ok {
	// 		deltas[k] = append(deltas[k], delta-prevDelta[0])
	// 	}
	// }
	// fmt.Println(deltas)

	// f(n) = f(n-11) + (860 + 162 * n) but this is only true for specific n starting at 64
	// can we use this to easily calculate the remaining values from X where X%11 is the same value as 26501365 (0), lets pick 66, so we have f(n) = f(n-11) + (880 + 162 * n)
	// idk because idk what the answer is for the input. but I get 470149643712804
	// lets just run ALL this logic again but for our real input and see what we get
	// -- Answer for the example --
	// newStepValues := map[int]int{}
	// newStepValues[66] = 2794
	// for k := 77; k <= 26501365; k += 11 {
	// 	newStepValues[k] = newStepValues[k-11] + 880 + (162 * ((k - 66) / 11))
	// }
	// fmt.Println(newStepValues[26501365])

	// annoyingly we have 26501365%131 = 65, so we'll need to start from 65 + (131*n), lets say step 327:97230, delta = 62148
	// f(N_i) = f(N_i-1) + (62148 + 31010*(i_1)) where N_i is in the sequence {327, 458...26501365}
	// IT WORKS!!!
	// not going to even try to code a general solution for this
	for k := 458; k <= 26501365; k += 131 {
		stepValues[k] = stepValues[k-131] + 62148 + (31010 * ((k - 327) / 131))
	}
	answer = stepValues[26501365]
	fmt.Printf("Part Two: %d\n", answer)
}

func getStartingCoord(input []string) u.Coord {
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				return u.NewCoord(x, y)
			}
		}
	}

	return u.NewCoord(0, 0)
}

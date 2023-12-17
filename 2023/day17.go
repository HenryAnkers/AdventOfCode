package main

import (
	u "adventofcode/utils"
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {
	inputFilePath := "./input/day17.txt"
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

type nodeState struct {
	steps     int
	direction int //0123 upleftdownright
	coord     u.Coord
}

func partOne(input []string) {
	// shortest path but with a twist? how do we manage the max 3 steps rule?
	// if we arrive at a node N on our 3rd step in a direction, we only have 2 neighbours. If we arrive at any other time we have 3...
	// can maybe consider (node, direction, numSteps) to decide if we've visited a node?
	// from this state we can calculate neighbour states and mark them as visited or not. it's probably a bit wasteful but there shouldnt be an unmanageable amount of states
	// end state is probably special case (if the coords equal the end coord we can return)
	answer := 0
	distances := map[nodeState]int{}
	visited := map[nodeState]bool{}
	startState1 := nodeState{
		steps:     1,
		coord:     u.NewCoord(1, 0),
		direction: 1,
	}
	startState2 := nodeState{
		steps:     1,
		coord:     u.NewCoord(0, 1),
		direction: 2,
	}
	nodeCost1, _ := strconv.Atoi(string(input[0][1]))
	nodeCost2, _ := strconv.Atoi(string(input[1][0]))

	distances[startState1] = nodeCost1
	distances[startState2] = nodeCost2
	queue := []nodeState{startState1, startState2}

search:
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			distanceI := distances[queue[i]]
			distanceJ := distances[queue[j]]

			return distanceI < distanceJ
		})

		nextNode := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []nodeState{}
		}

		visited[nextNode] = true
		newNeighbours := []nodeState{}
		newNeighbours = append(newNeighbours, nodeState{
			direction: (nextNode.direction + 1) % 4,
			steps:     1,
			coord:     nextNode.coord.NewCoordInDirection((nextNode.direction + 1) % 4),
		})
		newNeighbours = append(newNeighbours, nodeState{
			direction: (nextNode.direction + 3) % 4,
			steps:     1,
			coord:     nextNode.coord.NewCoordInDirection((nextNode.direction + 3) % 4),
		})
		if nextNode.steps < 3 {
			newNeighbours = append(newNeighbours, nodeState{
				direction: nextNode.direction,
				steps:     nextNode.steps + 1,
				coord:     nextNode.coord.NewCoordInDirection(nextNode.direction),
			})
		}

		for _, node := range newNeighbours {
			if node.coord.Y < 0 || node.coord.Y >= len(input) || node.coord.X < 0 || node.coord.X >= len(input[node.coord.Y]) || visited[node] {
				continue
			}

			nodeCost, _ := strconv.Atoi(string(input[node.coord.Y][node.coord.X]))
			if node.coord.Y == len(input)-1 && node.coord.X == len(input[node.coord.Y])-1 {
				answer = distances[nextNode] + nodeCost
				break search
			}

			if oldDistance, ok := distances[node]; ok {
				distances[node] = min(oldDistance, distances[nextNode]+nodeCost)
			} else {
				distances[node] = distances[nextNode] + nodeCost
				queue = append(queue, node)
			}
		}
	}

	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// already 30s runtime for part 1 D:
	// im sure replacing queue with better DS (min heap?) would help a lot but I dont think part 2 should be THAT much worse...
	// in the end ~6mins using queue slice, which is quicker than me implementing a min heap
	// improvement with a min heap? 352s -> 1.91s
	answer := 0

	distances := map[nodeState]int{}
	visited := map[nodeState]bool{}
	startState1 := nodeState{
		steps:     1,
		coord:     u.NewCoord(1, 0),
		direction: 1,
	}
	startState2 := nodeState{
		steps:     1,
		coord:     u.NewCoord(0, 1),
		direction: 2,
	}
	nodeCost1, _ := strconv.Atoi(string(input[0][1]))
	nodeCost2, _ := strconv.Atoi(string(input[1][0]))

	distances[startState1] = nodeCost1
	distances[startState2] = nodeCost2
	queue := []nodeState{startState1, startState2}

	priorityQueue := &PriorityQueue{
		items:     queue,
		distances: distances,
	}
	heap.Init(priorityQueue)

search:
	for priorityQueue.Len() > 0 {
		nextNode := heap.Pop(priorityQueue).(nodeState)

		visited[nextNode] = true
		newNeighbours := []nodeState{}

		if nextNode.steps < 4 {
			newNeighbours = append(newNeighbours, nodeState{
				direction: nextNode.direction,
				steps:     nextNode.steps + 1,
				coord:     nextNode.coord.NewCoordInDirection(nextNode.direction),
			})
		} else {
			newNeighbours = append(newNeighbours, nodeState{
				direction: (nextNode.direction + 1) % 4,
				steps:     1,
				coord:     nextNode.coord.NewCoordInDirection((nextNode.direction + 1) % 4),
			})
			newNeighbours = append(newNeighbours, nodeState{
				direction: (nextNode.direction + 3) % 4,
				steps:     1,
				coord:     nextNode.coord.NewCoordInDirection((nextNode.direction + 3) % 4),
			})
			if nextNode.steps < 10 {
				newNeighbours = append(newNeighbours, nodeState{
					direction: nextNode.direction,
					steps:     nextNode.steps + 1,
					coord:     nextNode.coord.NewCoordInDirection(nextNode.direction),
				})
			}
		}

		for _, node := range newNeighbours {
			if node.coord.Y < 0 || node.coord.Y >= len(input) || node.coord.X < 0 || node.coord.X >= len(input[node.coord.Y]) || visited[node] {
				continue
			}

			nodeCost, _ := strconv.Atoi(string(input[node.coord.Y][node.coord.X]))
			// this caught me out, need to make sure we've already taken at least 4 steps to reach this node before considering it valid...
			if node.coord.Y == len(input)-1 && node.coord.X == len(input[node.coord.Y])-1 && node.steps >= 4 {
				answer = distances[nextNode] + nodeCost
				break search
			}

			if oldDistance, ok := distances[node]; ok {
				distances[node] = min(oldDistance, distances[nextNode]+nodeCost)
			} else {
				distances[node] = distances[nextNode] + nodeCost
				heap.Push(priorityQueue, node)
			}
		}
	}

	fmt.Printf("Part Two: %d\n", answer)
}

type PriorityQueue struct {
	items     []nodeState
	distances map[nodeState]int
}

func (pq PriorityQueue) Len() int { return len(pq.items) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.distances[pq.items[i]] < pq.distances[pq.items[j]]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.items = append(pq.items, x.(nodeState))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	x := old[n-1]
	pq.items = old[0 : n-1]
	return x
}

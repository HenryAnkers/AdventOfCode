package main

import (
	"fmt"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day20.txt"
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

type Node struct {
	nodeType       string
	isOn           bool
	previousInputs map[string]bool //string is node name, we will use 0 for low, 1 for high pulses (isHigh)
	childrenNodes  []string
}

type Pulse struct {
	source      string
	destination string
	isHigh      bool
}

func partOne(input []string) {
	// looks interesting... I bet part 2 is what happens after we push a billion times or something
	// for FlipFlop we have State(on,off), childrenNodes[] string
	// for Conjunction we have PrevInputs(map[parentNode]pulseType), childrenNodes []string
	// will just use same type to represent both
	answer := 0
	numLow := 0
	numHigh := 0
	broadcastNodes, nodes := parseNodes(input)
	initialQueue := []Pulse{}
	for _, name := range broadcastNodes {
		initialQueue = append(initialQueue, Pulse{
			source:      "broadcast",
			destination: name,
			isHigh:      false,
		})
	}

	for i := 0; i < 1000; i++ {
		queue := []Pulse{}
		queue = append(queue, initialQueue...)
		numLow += (1 + len(initialQueue))

		for len(queue) > 0 {
			pulseToProcess := queue[0]
			queue = queue[1:]
			newPulses := runPulse(pulseToProcess, nodes)

			for _, p := range newPulses {
				if p.isHigh {
					numHigh++
				} else {
					numLow++
				}
			}

			queue = append(queue, newPulses...)
		}
	}

	answer = numHigh * numLow
	fmt.Printf("Part One: %d\n", answer)
}

func partTwo(input []string) {
	// is this as easy as it seems? just iterate until we find it? surely not?
	// whilst it runs... there must be a none bruteforce solution to this
	// must have something to do with cycles, but given its the first time RX gets hit it's a bit confusing. does some part of the graph repeat? can we skip some (but not all) of the states?
	// use findSubgraphs on the broadcast node, can see each one sort of splits into different graphs with some overlapping at the end. actually only 2 nodes overlap and theyre in all 4 sets
	// this means theres probably an independent cycle in each one, then the zr/rx ones only get triggered at some certain point?
	// if all cycle lengths were disjoint (e.g. 5/7/13) then we might need to wait for the LCM of all of them? hard to know exactly what we're waiting for though without examining the graph further (how do we know its not the 100th time they intersect or something)
	// how to find out... can we just run each starting node individually until we find a cycle (use a memo of the states of all the nodes)? we can then maybe examine what the two overlapping node states are at the end of this (and then just skip the LCM of those cycle steps and repeat)
	// NVM ITS JUST THE LCM OF ALL THESE CYCLES!!! dont know why and dont care
	// it looks like zr is a Conjunction type, probably sends to rx when receives input from each of its children, and that happens exactly once at the end of their cycles.
	answer := 1

	broadcastNodes, nodes := parseNodes(input)
	initialQueue := []Pulse{}
	for _, name := range broadcastNodes {
		initialQueue = append(initialQueue, Pulse{
			source:      "broadcast",
			destination: name,
			isHigh:      false,
		})
	}

	//findSubgraphs(broadcastNodes, nodes)
	// nodesToProcess is taken from output of findSubgraphs pasted into chatGPT with the overlapping nodes being removed (theres exactly 2 in each)
	nodesToProcess := map[int][]string{
		0: []string{"bj", "jl", "mn", "ps", "pz", "tc", "td", "vl", "vz", "xf", "xg", "xs", "xx", "zk"},
		1: []string{"bk", "cm", "cs", "gb", "jn", "km", "ks", "mf", "ml", "qx", "rh", "vd", "xk", "zz"},
		2: []string{"dn", "dz", "gc", "hd", "hf", "jk", "kf", "mm", "nr", "nv", "pn", "qf", "zb", "zf"},
		3: []string{"cp", "fl", "fm", "gx", "jm", "kb", "lq", "ms", "mz", "nz", "pp", "sz", "tl", "zv"},
	}
	allCycleLengths := []int{}

	for i, ip := range initialQueue {
		memo := map[string]int{}
		startQueue := []Pulse{ip}
		pressed := 0
		_, nodes = parseNodes(input)
		currentNodeList := nodesToProcess[i]

	processing:
		for true {
			pressed += 1
			key := nodeStateKey(nodes, currentNodeList)
			if prev, ok := memo[key]; ok {
				allCycleLengths = append(allCycleLengths, pressed-prev)
				break processing
			} else {
				memo[key] = pressed
			}

			queue := []Pulse{}
			queue = append(queue, startQueue...)

			for len(queue) > 0 {
				pulseToProcess := queue[0]
				queue = queue[1:]
				newPulses := runPulse(pulseToProcess, nodes)
				queue = append(queue, newPulses...)
			}
		}

	}

	for _, length := range allCycleLengths {
		answer = utils.LCM(answer, length)
	}

	fmt.Printf("Part Two: %d\n", answer)
}

func nodeStateKey(nodes map[string]*Node, nodesToInclude []string) string {
	out := ""

	for _, name := range nodesToInclude {
		if node, ok := nodes[name]; ok {
			state := "off"
			if node.isOn {
				state = "high"
			}

			childInputString := ""
			for parentName, previousInputWasHigh := range node.previousInputs {
				inputType := "low"
				if previousInputWasHigh {
					inputType = "high"
				}
				childInputString += "(" + parentName + "__" + inputType + ")"
			}

			out += ("{" + name + "--" + childInputString + "--" + state + "} ")
		}
	}

	return out
}
func parseNodes(input []string) ([]string, map[string]*Node) {
	broadcastNodes := []string{}
	nodes := map[string]*Node{}

	for _, line := range input {
		nodeName := strings.Split(line, " -> ")[0]
		destinations := strings.Split(strings.Split(line, " -> ")[1], ", ")

		if nodeName == "broadcaster" {
			broadcastNodes = destinations
		} else {
			nodeType := nodeName[0]
			name := nodeName[1:]

			if nodeType == '%' {
				nodes[name] = &Node{
					nodeType:      "flip",
					isOn:          false,
					childrenNodes: destinations,
				}
			} else {
				nodes[name] = &Node{
					nodeType:       "conj",
					previousInputs: map[string]bool{},
					childrenNodes:  destinations,
				}
			}
		}
	}

	for parentName, node := range nodes {
		for _, childNodeName := range node.childrenNodes {
			if childNode, ok := nodes[childNodeName]; ok && childNode.nodeType == "conj" {
				childNode.previousInputs[parentName] = false
			}
		}
	}

	return broadcastNodes, nodes
}

func runPulse(pulse Pulse, nodes map[string]*Node) []Pulse {
	nodeName := pulse.destination
	newPulses := []Pulse{}
	nodeHit, ok := nodes[nodeName]
	if !ok {
		return newPulses
	}

	if nodeHit.nodeType == "flip" && !pulse.isHigh {
		if nodeHit.isOn { // send low, turn off
			for _, name := range nodeHit.childrenNodes {
				newPulses = append(newPulses, Pulse{
					source:      nodeName,
					destination: name,
					isHigh:      false,
				})
			}
			nodeHit.isOn = false
		} else { //send high, turn on
			for _, name := range nodeHit.childrenNodes {
				newPulses = append(newPulses, Pulse{
					source:      nodeName,
					destination: name,
					isHigh:      true,
				})
			}
			nodeHit.isOn = true
		}
	} else if nodeHit.nodeType == "conj" {
		nodeHit.previousInputs[pulse.source] = pulse.isHigh

		resultingPulseIsHigh := false
		for _, wasHigh := range nodeHit.previousInputs {
			if !wasHigh {
				resultingPulseIsHigh = true
				break
			}
		}

		for _, name := range nodeHit.childrenNodes {
			newPulses = append(newPulses, Pulse{
				source:      nodeName,
				destination: name,
				isHigh:      resultingPulseIsHigh,
			})
		}
	}

	return newPulses
}

func findSubgraphs(inputNodes []string, nodes map[string]*Node) {
	for _, node := range inputNodes {
		visited := map[string]bool{}
		queue := []string{node}
		for len(queue) > 0 {
			newQueue := []string{}
			for _, nodeName := range queue {
				if !visited[nodeName] {
					if n, ok := nodes[nodeName]; ok {
						newQueue = append(newQueue, n.childrenNodes...)
					}
					visited[nodeName] = true
				}
			}
			queue = newQueue
		}
		fmt.Println(visited)
	}
}

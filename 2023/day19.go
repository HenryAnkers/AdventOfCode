package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFilePath := "./input/day19.txt"
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	start := time.Now()
	partOne(string(input))
	duration1 := time.Since(start)

	start = time.Now()
	partTwo(string(input))
	duration2 := time.Since(start)
	fmt.Printf("Part One took: %v, Part Two took: %v\n", duration1, duration2)
}

type Part struct {
	x int
	m int
	a int
	s int
}

func partOne(input string) {
	// dont really like this...
	// parse the input into a list of rules
	// each rule has a character (x,m,a,s) and a destination. character = nil => true by default
	// save list of rules in a map where the key string is the name of rule... value is list of rules, iterate through them in order
	answer := 0

	ruleBlockStrings := strings.Split(strings.Split(input, "\r\n\r\n")[0], "\r\n")
	partBlockStrings := strings.Split(strings.Split(input, "\r\n\r\n")[1], "\r\n")
	ruleMap := map[string][]string{} // will just parse the rules on the fly
	for _, rule := range ruleBlockStrings {
		ruleName := strings.Split(rule, "{")[0]
		ruleList := strings.Split(strings.Split(strings.Split(rule, "{")[1], "}")[0], ",")
		ruleMap[ruleName] = ruleList
	}
	for _, part := range partBlockStrings {
		currRuleName := "in"
		currentPart := Part{}
		partList := strings.Split(strings.Split(strings.Split(part, "{")[1], "}")[0], ",")
		currentPart.x, _ = strconv.Atoi(strings.Split(partList[0], "=")[1])
		currentPart.m, _ = strconv.Atoi(strings.Split(partList[1], "=")[1])
		currentPart.a, _ = strconv.Atoi(strings.Split(partList[2], "=")[1])
		currentPart.s, _ = strconv.Atoi(strings.Split(partList[3], "=")[1])

		for currRuleName != "A" && currRuleName != "R" {
			currRule := ruleMap[currRuleName]
			for _, rule := range currRule {
				if strings.Contains(rule, ":") {
					criteria := strings.Split(rule, ":")[0]
					criteriaAmount, _ := strconv.Atoi(criteria[2:])
					destination := strings.Split(rule, ":")[1]
					metCriteria := false

					ruleKey := rule[0:2]
					switch ruleKey {
					case "x<":
						metCriteria = currentPart.x < criteriaAmount
					case "x>":
						metCriteria = currentPart.x > criteriaAmount
					case "m<":
						metCriteria = currentPart.m < criteriaAmount
					case "m>":
						metCriteria = currentPart.m > criteriaAmount
					case "a<":
						metCriteria = currentPart.a < criteriaAmount
					case "a>":
						metCriteria = currentPart.a > criteriaAmount
					case "s<":
						metCriteria = currentPart.s < criteriaAmount
					case "s>":
						metCriteria = currentPart.s > criteriaAmount
					}

					if metCriteria {
						currRuleName = destination
						break
					}
				} else {
					currRuleName = rule
					break
				}
			}
		}

		if currRuleName == "A" {
			answer += currentPart.x
			answer += currentPart.m
			answer += currentPart.a
			answer += currentPart.s
		}

	}

	fmt.Printf("Part One: %d\n", answer)
}

type Part2 struct {
	x []int
	m []int
	a []int
	s []int
}

func partTwo(input string) {
	// hmmm
	// i wonder if we can use some sort of DP (i dont know the actual name)
	// start with X,M,A,S in the range [1,4000], rule = in
	// we then split into having (X,M,A = [1,4000], S = [1,1350], rule = px) and (X,M,A = [1,4000], S = [1351,4000], rule = qqz)
	// when we hit an A we add all the ranges to valid, when we hit R we just delete them all
	// we then merge all the ranges together at the end and multiply (theyre inclusive of first and last number)
	answer := 0

	ruleBlockStrings := strings.Split(strings.Split(input, "\r\n\r\n")[0], "\r\n")
	ruleMap := map[string][]string{}
	for _, rule := range ruleBlockStrings {
		ruleName := strings.Split(rule, "{")[0]
		ruleList := strings.Split(strings.Split(strings.Split(rule, "{")[1], "}")[0], ",")
		ruleMap[ruleName] = ruleList
	}

	initialRanges := Part2{
		x: []int{1, 4000},
		m: []int{1, 4000},
		a: []int{1, 4000},
		s: []int{1, 4000},
	}
	initialRuleName := "in"
	allValidRanges := getValidRanges(initialRanges, initialRuleName, ruleMap)
	for _, part := range allValidRanges {
		answer += ((part.x[1] - part.x[0] + 1) * (part.m[1] - part.m[0] + 1) * (part.a[1] - part.a[0] + 1) * (part.s[1] - part.s[0] + 1))
	}

	// cannot believe this worked first time...
	fmt.Printf("Part Two: %d\n", answer)
}

func getValidRanges(initialPart Part2, currentRuleName string, ruleMap map[string][]string) []Part2 {
	if currentRuleName == "A" {
		return []Part2{initialPart}
	}
	if currentRuleName == "R" {
		return []Part2{}
	}

	ruleList := ruleMap[currentRuleName]
	results := []Part2{}
	part := initialPart
	for _, rule := range ruleList {
		if strings.Contains(rule, ":") {
			criteria := strings.Split(rule, ":")[0]
			criteriaAmount, _ := strconv.Atoi(criteria[2:])
			destination := strings.Split(rule, ":")[1]

			ruleKey := rule[0:2]
			switch ruleKey {
			// splitting logic is awkward
			// if the whole range satisfies the rule... then we only need to consider this path
			// if part of the range satisfies the rule...
			// - we need to split into the values of X that satisfy the rule in a new Part2 object and continue down that path
			// - we then need to modify the part we're processing (create a new part and point to it, lets not mutate the parts or individual ranges) to only include values of X that dont satisfy and continue
			// if none then we continue on similar to above but our current part is unchanged
			// sorry for the crap copy and paste job below
			case "x<":
				if part.x[1] < criteriaAmount {
					// case 1: the whole range is valid
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.x[0] < criteriaAmount {
					// case 2: at least one number is valid and caught by this rule
					newRange1 := Part2{
						x: []int{part.x[0], criteriaAmount - 1},
						m: part.m,
						a: part.a,
						s: part.s,
					}
					part = Part2{ // we must now continue on with this part...
						x: []int{criteriaAmount, part.x[1]},
						m: part.m,
						a: part.a,
						s: part.s,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}

			case "x>":
				if part.x[0] > criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.x[1] > criteriaAmount {
					newRange1 := Part2{
						x: []int{criteriaAmount + 1, part.x[1]},
						m: part.m,
						a: part.a,
						s: part.s,
					}
					part = Part2{
						x: []int{part.x[0], criteriaAmount},
						m: part.m,
						a: part.a,
						s: part.s,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}
			case "m<":
				if part.m[1] < criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.m[0] < criteriaAmount {
					newRange1 := Part2{
						m: []int{part.m[0], criteriaAmount - 1},
						x: part.x,
						a: part.a,
						s: part.s,
					}
					part = Part2{
						m: []int{criteriaAmount, part.m[1]},
						x: part.x,
						a: part.a,
						s: part.s,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}

			case "m>":
				if part.m[0] > criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.m[1] > criteriaAmount {
					newRange1 := Part2{
						m: []int{criteriaAmount + 1, part.m[1]},
						x: part.x,
						a: part.a,
						s: part.s,
					}
					part = Part2{
						m: []int{part.m[0], criteriaAmount},
						x: part.x,
						a: part.a,
						s: part.s,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}

			case "a<":
				if part.a[1] < criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.a[0] < criteriaAmount {
					newRange1 := Part2{
						a: []int{part.a[0], criteriaAmount - 1},
						x: part.x,
						m: part.m,
						s: part.s,
					}
					part = Part2{
						a: []int{criteriaAmount, part.a[1]},
						x: part.x,
						m: part.m,
						s: part.s,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}

			case "a>":
				if part.a[0] > criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.a[1] > criteriaAmount {
					newRange1 := Part2{
						a: []int{criteriaAmount + 1, part.a[1]},
						x: part.x,
						m: part.m,
						s: part.s,
					}
					part = Part2{
						a: []int{part.a[0], criteriaAmount},
						x: part.x,
						m: part.m,
						s: part.s,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}

			case "s<":
				if part.s[1] < criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.s[0] < criteriaAmount {
					newRange1 := Part2{
						s: []int{part.s[0], criteriaAmount - 1},
						x: part.x,
						m: part.m,
						a: part.a,
					}
					part = Part2{
						s: []int{criteriaAmount, part.s[1]},
						x: part.x,
						m: part.m,
						a: part.a,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}

			case "s>":
				if part.s[0] > criteriaAmount {
					return append(results, getValidRanges(part, destination, ruleMap)...)
				} else if part.s[1] > criteriaAmount {
					newRange1 := Part2{
						s: []int{criteriaAmount + 1, part.s[1]},
						x: part.x,
						m: part.m,
						a: part.a,
					}
					part = Part2{
						s: []int{part.s[0], criteriaAmount},
						x: part.x,
						m: part.m,
						a: part.a,
					}
					results = append(results, getValidRanges(newRange1, destination, ruleMap)...)
				}
			}
		} else {
			return append(results, getValidRanges(part, rule, ruleMap)...)
		}
	}
	return results
}

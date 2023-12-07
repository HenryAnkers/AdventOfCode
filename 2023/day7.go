package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day7.txt"
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
	answer := 0
	handBidMapping := map[string]int{}
	allHands := []string{}
	cardStrengths := map[string]int{
		"A": 12,
		"K": 11,
		"Q": 10,
		"J": 9,
		"T": 8,
		"9": 7,
		"8": 6,
		"7": 5,
		"6": 4,
		"5": 3,
		"4": 2,
		"3": 1,
		"2": 0,
	}

	for _, line := range input {
		hand := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])

		allHands = append(allHands, hand)
		handBidMapping[hand] = bid
	}

	sort.Slice(allHands, func(i, j int) bool {
		handI := allHands[i]
		handJ := allHands[j]
		typeI := getTypePartOne(handI)
		typeJ := getTypePartOne(handJ)

		if typeJ > typeI {
			return true
		} else if typeJ < typeI {
			return false
		}

		for p := 0; p < 5; p++ {
			if cardStrengths[string(handJ[p])] > cardStrengths[string(handI[p])] {
				return true
			} else if cardStrengths[string(handJ[p])] < cardStrengths[string(handI[p])] {
				return false
			}
		}

		return true
	})

	for i, hand := range allHands {
		answer += (i + 1) * handBidMapping[hand]
	}

	fmt.Printf("Part One: %d\n", answer)
}

// 0 = High card, 1 = One pair, 2 = Two pair, 3 = Three of a kind, 4 = Full house, 5 = Four of a kind, 6 = Five of a kind
func getTypePartOne(hand string) int {
	cardCount := map[string]int{}

	for _, card := range hand {
		cardCount[string(card)] += 1
	}

	numPairs := 0
	hasThree := false

	for _, amount := range cardCount {
		if amount == 5 || amount == 4 {
			// Five of a kind, four of a kind
			return amount + 1
		}
		if amount == 3 {
			hasThree = true
		} else if amount == 2 {
			numPairs += 1
		}
	}

	if hasThree && numPairs == 1 {
		// Full house
		return 4
	} else if hasThree {
		// Three of a kind
		return 3
	} else if numPairs > 0 {
		//1 or 2, representing 1 pair 2 pair respectively
		return numPairs
	} else {
		//High card
		return 0
	}
}

func partTwo(input []string) {
	answer := 0

	handBidMapping := map[string]int{}
	allHands := []string{}
	cardStrengths := map[string]int{
		"A": 12,
		"K": 11,
		"Q": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
		"J": 0,
	}

	for _, line := range input {
		hand := strings.Fields(line)[0]
		bid, _ := strconv.Atoi(strings.Fields(line)[1])

		allHands = append(allHands, hand)
		handBidMapping[hand] = bid
	}

	sort.Slice(allHands, func(i, j int) bool {
		handI := allHands[i]
		handJ := allHands[j]
		typeI := getTypePartTwo(handI)
		typeJ := getTypePartTwo(handJ)

		if typeJ > typeI {
			return true
		} else if typeJ < typeI {
			return false
		}

		for p := 0; p < 5; p++ {
			if cardStrengths[string(handJ[p])] > cardStrengths[string(handI[p])] {
				return true
			} else if cardStrengths[string(handJ[p])] < cardStrengths[string(handI[p])] {
				return false
			}
		}

		return true
	})

	for i, hand := range allHands {
		answer += (i + 1) * handBidMapping[hand]
	}

	fmt.Printf("Part Two: %d\n", answer)
}

// 0 = High card, 1 = One pair, 2 = Two pair, 3 = Three of a kind, 4 = Full house, 5 = Four of a kind, 6 = Five of a kind
func getTypePartTwo(hand string) int {
	cardCount := map[string]int{}

	numJok := 0
	for _, card := range hand {
		cardCount[string(card)] += 1
		if string(card) == "J" {
			numJok += 1
		}
	}

	numPairs := 0
	hasThree := false

	cards := strings.Split(hand, "")
	sort.Slice(cards, func(i, j int) bool {
		return cardCount[cards[i]] > cardCount[cards[j]]
	})

	processedCard := map[string]bool{}

	//rough strategy - use the joker at the first opportunity on the card with the current highest occurance. dont consider joker itself as a valid card (unless its JJJJJ...)
	for _, card := range cards {
		if processedCard[card] == true {
			continue
		}

		amount := cardCount[card]
		if amount == 5 || numJok == 5 || amount+numJok == 5 {
			return 6
		} else if string(card) == "J" {
			continue
		} else if amount == 4 || amount+numJok == 4 {
			return 5
		} else if amount == 3 {
			hasThree = true
		} else if amount+numJok == 3 {
			hasThree = true
			numJok = 0
		} else if amount == 2 {
			numPairs += 1
		} else if amount+numJok == 2 {
			numPairs += 1
			numJok = 0
		}

		processedCard[card] = true
	}

	if hasThree && numPairs == 1 {
		// Full house
		return 4
	} else if hasThree {
		// Three of a kind
		return 3
	} else if numPairs > 0 {
		//1 or 2, representing 1 pair 2 pair respectively
		return numPairs
	} else {
		//High card
		return 0
	}
}

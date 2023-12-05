package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"adventofcode/utils"
)

func main() {
	inputFilePath := "./input/day5.txt"
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

// So we need to find any offsets in the seed to soil map by parsing the rows
// The real inputs have huge ranges and there are a lot of them. Building a map of the individual offsets is a bad idea. We'll save the ranges instead in a new struct and call a helper function to use them
// input of a b c -> for items in range(b, b+c), the destination for source x is equal to x + (a-b)
// We do this, start with our initial array of seed values, build them step by step into location values, done
// seed -> soil -> fertilizer -> water -> light -> temperature - humidity

type mapValue struct {
	lowInt  int // the int that the source starts at
	highInt int // the maximum int the range extends to
	offset  int // the offset we apply to each source item x where lowint <= x <= highint to get the destination
}

type itemRange struct {
	low  int
	high int
}

func partOne(input []string) {
	answer := 0

	seeds, itemMaps := parseInput(input)

	currentInputArray := seeds
	for _, itemMap := range itemMaps {
		newInputArray := []int{}

		for _, item := range currentInputArray {
			//check if any of the values in the map applies to the seed, if so add the offset
			itemOutput := item

			for _, mapRange := range itemMap {
				if item >= mapRange.lowInt && item <= mapRange.highInt {
					itemOutput += mapRange.offset
					break // it's not given that seeds can only be included in one range but running without break gives same answer
				}
			}

			newInputArray = append(newInputArray, itemOutput)
		}
		currentInputArray = newInputArray
	}

	answer = slices.Min(currentInputArray)

	fmt.Printf("Part One: %d\n", answer)
}

// this isnt really multiplication, it's a weird form of finding the intersection and adding the offset to it
func multiplyRangeByMapValue(itemranges []itemRange, mapvalue mapValue) ([]itemRange, []itemRange) {
	// range = (a,b), mapvalue = (c,d), offset = o

	// keep track of which ones might need further processing and which can't be...
	existingItemRanges := []itemRange{}
	newItemRanges := []itemRange{}

	for _, itemrange := range itemranges {
		initialRangeStart := itemrange.low
		initialRangeEnd := itemrange.high
		mapValueStart := mapvalue.lowInt
		mapValueEnd := mapvalue.highInt
		offset := mapvalue.offset

		if initialRangeStart < mapValueStart && initialRangeEnd > mapValueEnd {
			// Three new ranges

			existingItemRanges = append(existingItemRanges, []itemRange{
				{low: initialRangeStart, high: mapValueStart - 1},
				{low: mapValueEnd + 1, high: initialRangeEnd},
			}...)
			newItemRanges = append(newItemRanges, []itemRange{
				{low: mapValueStart + offset, high: mapValueEnd + offset},
			}...)
		} else if initialRangeStart >= mapValueStart && initialRangeEnd > mapValueEnd && initialRangeStart < mapValueEnd {
			// Two new ranges, overlap at the start
			existingItemRanges = append(existingItemRanges, []itemRange{
				{low: mapValueEnd + 1, high: initialRangeEnd},
			}...)
			newItemRanges = append(newItemRanges, []itemRange{
				{low: initialRangeStart + offset, high: mapValueEnd + offset},
			}...)
		} else if initialRangeStart < mapValueStart && initialRangeEnd <= mapValueEnd && mapValueStart < initialRangeEnd {
			// Two new ranges, overlap at the end
			existingItemRanges = append(existingItemRanges, []itemRange{
				{low: initialRangeStart, high: mapValueStart - 1},
			}...)
			newItemRanges = append(newItemRanges, []itemRange{
				{low: mapValueStart + offset, high: initialRangeEnd + offset},
			}...)
		} else if initialRangeStart >= mapValueStart && initialRangeEnd <= mapValueEnd {
			// Entire range is overlapped
			newItemRanges = append(newItemRanges, []itemRange{
				{low: initialRangeStart + offset, high: initialRangeEnd + offset},
			}...)
		} else { // no overlap at all, this might still need processing
			existingItemRanges = append(existingItemRanges, itemrange)
		}

	}

	return existingItemRanges, newItemRanges
}

func partTwo(input []string) {
	// Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.
	// > ranges of seed numbers
	// is there an efficient way of doing this or do we really need to simulate all these seeds?
	// note: brute force solution gievs me OOM errors so not an option
	// given a range of seeds (a,b), we know that our current itemMaps (low,high,offset) will apply to another range of them.... this basically gives us up to three new ranges.
	// so  we can be efficient here if we define an operation between (a,b) and (low, high, offset) to get another list of ranges..
	// we would then need to squash the ranges down.. and our answer would be the lowest value a_i for (a_n, b_n)....
	// if we dont squash the ranges we get 3^6 * original number which isn't bad and maybe doable...
	// feels overtly complicated

	answer := 0

	fakeSeeds, itemMaps := parseInput(input)
	itemRanges := []itemRange{}
	for i := 0; i < len(fakeSeeds); i += 2 {
		lowValue := fakeSeeds[i]
		amount := fakeSeeds[i+1]
		itemRanges = append(itemRanges, itemRange{low: lowValue, high: lowValue + amount - 1})
	}

	for _, itemMap := range itemMaps {
		newItemRanges := []itemRange{}

		// ok on reflection this is wrong
		// one of our map items might intersect the first part of the range, another the second part... but they need to be dealt with at the same time, not seperately
		// pass in a list of ranges, get out a list of ranges...

		// we need to apply any future operations to only new ranges, not ones that have been created by applying an offset, ffs
		for _, itemrange := range itemRanges {
			currentItemRangeList := []itemRange{itemrange}

			for _, mapRange := range itemMap {
				var newItemRangeList []itemRange
				currentItemRangeList, newItemRangeList = multiplyRangeByMapValue(currentItemRangeList, mapRange)

				// we won't attempt to process these ranges further, just save them for next time
				// if you are trying to read this code I am so sorry about these variable names
				newItemRanges = append(newItemRanges, newItemRangeList...)
			}

			// save remaining unprocessed ranges
			newItemRanges = append(newItemRanges, currentItemRangeList...)
		}

		itemRanges = newItemRanges
	}

	answer = math.MaxInt
	for _, r := range itemRanges {
		answer = min(answer, r.low)
	}

	fmt.Printf("Part Two: %d\n", answer)
}

func parseInput(input []string) ([]int, [][]mapValue) {
	var seeds []int
	var itemMaps [][]mapValue
	var currentMap []mapValue

	for _, line := range input {
		if strings.Contains(line, ":") {
			if len(currentMap) > 0 {
				itemMaps = append(itemMaps, currentMap)
				currentMap = nil
			}
			if strings.Contains(line, "seeds") {
				seedStrings := strings.Fields(strings.Split(line, ":")[1])
				for _, v := range seedStrings {
					seedValue, _ := strconv.Atoi(v)
					seeds = append(seeds, seedValue)
				}
			}
		} else if line != "" {
			valueStringList := strings.Fields(line)
			destinationValue, _ := strconv.Atoi(valueStringList[0])
			sourceValue, _ := strconv.Atoi(valueStringList[1])
			rangeValue, _ := strconv.Atoi(valueStringList[2])

			newItemMapping := mapValue{
				lowInt:  sourceValue,
				highInt: sourceValue + rangeValue - 1, // remember it's inclusive
				offset:  destinationValue - sourceValue,
			}

			currentMap = append(currentMap, newItemMapping)
		}
	}

	if len(currentMap) > 0 {
		itemMaps = append(itemMaps, currentMap)
	}

	return seeds, itemMaps
}

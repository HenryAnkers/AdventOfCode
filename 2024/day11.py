import os
import time
from collections import Counter
from collections import defaultdict

def parseIntInput(file_path):
    with open(file_path, 'r') as file:
        return [[int(item) for item in line.strip().split()] for line in file]

def part1(data):
    stones = data
    for _ in range(0, 25):
        newStones = []
        for stone in stones:
            if stone == 0:
                newStones.append(1)
            elif len(str(stone)) % 2 == 0:
                stoneStr = str(stone)
                length = len(stoneStr) // 2
                newStones.append(int(stoneStr[:length]))
                newStones.append(int(stoneStr[length:]))
            else:
                newStones.append(stone * 2024)
        stones = newStones

    return len(stones)

def part2(data):
    # we clearly get lots of duplicate stones, lets try using dictionary instead?
    stoneCount = Counter(data)
    for _ in range(0, 75):
        newStoneCount = defaultdict(int)
        for stone in stoneCount:
            newStones = []
            if stone == 0:
                newStones.append(1)
            elif len(str(stone)) % 2 == 0:
                stoneStr = str(stone)
                length = len(stoneStr) // 2
                newStones.append(int(stoneStr[:length]))
                newStones.append(int(stoneStr[length:]))
            else:
                newStones.append(stone * 2024)   

            for newStone in newStones:
                newStoneCount[newStone] += stoneCount[stone]
        stoneCount = newStoneCount

    return sum([stoneCount[c] for c in stoneCount])
   

day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseIntInput(input_path)[0]
end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")
print(input_data)

start_time = time.time()
solution1 = part1(input_data)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")
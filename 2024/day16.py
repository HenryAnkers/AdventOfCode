import os
import time
import sys
from collections import defaultdict

sys.setrecursionlimit(5000)


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def parseIntInput(file_path):
    with open(file_path, "r") as file:
        return [[int(item) for item in line.strip().split()] for line in file]


def parseInputAsMap(file_path):
    with open(file_path, "r") as file:
        data = [line.strip() for line in file]
        dataMap = defaultdict(lambda: defaultdict(bool))
        for y, line in enumerate(data):
            for x, char in enumerate(line):
                dataMap[(x, y)] = char
        return dataMap


def findPaths(dataMap, position, dIndex, score, memo):
    if position in memo and memo[position] < score - 1001:
        return -1, set()
    memo[position] = score

    if dataMap[position] == "E":
        return score, set([position])

    if dataMap[position] == "#":
        return -1, set()

    deltas = [(1, 0), (0, -1), (-1, 0), (0, 1)]
    x, y = position[0], position[1]

    # current direction
    dX, dY = deltas[dIndex][0], deltas[dIndex][1]
    move1, path1 = findPaths(dataMap, (x + dX, y + dY), dIndex, score + 1, memo)

    # 90 degrees
    newD = (dIndex - 1) % 4
    dX, dY = deltas[newD][0], deltas[newD][1]
    move2, path2 = findPaths(dataMap, (x + dX, y + dY), newD, score + 1001, memo)

    # 270 degrees
    newD = (dIndex + 1) % 4
    dX, dY = deltas[newD][0], deltas[newD][1]
    move3, path3 = findPaths(dataMap, (x + dX, y + dY), newD, score + 1001, memo)

    validScoresAndPaths = [s for s in [[move1, path1], [move2, path2], [move3, path3]] if s[0] != -1]
    minScore = min([s[0] for s in validScoresAndPaths]) if len(validScoresAndPaths) > 0 else -1
    if minScore == -1:
        return -1, set()

    validPathTiles = set([position])
    for x in validScoresAndPaths:
        if x[0] == minScore:
            validPathTiles = validPathTiles | x[1]

    return (minScore, validPathTiles)


def part1(data):
    maxY = max(p[1] for p in data.keys())
    startPos = (1, maxY - 1)
    return findPaths(data, startPos, 0, 0, {})[0]


def part2(data):
    maxY = max(p[1] for p in data.keys())
    startPos = (1, maxY - 1)
    return len(findPaths(data, startPos, 0, 0, {})[1])


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInputAsMap(input_path)
end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")

start_time = time.time()
solution1 = part1(input_data)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")

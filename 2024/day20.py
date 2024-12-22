import os
import time
import sys
from collections import defaultdict
from collections import Counter

sys.setrecursionlimit(10000)


def parseInputAsMap(file_path):
    with open(file_path, "r") as file:
        data = [line.strip() for line in file]
        dataMap = defaultdict(bool)
        startPos = (0, 0)
        endPo = (0, 0)
        for y, line in enumerate(data):
            for x, char in enumerate(line):
                dataMap[(x, y)] = char
                if char == "S":
                    startPos = (x, y)
                elif char == "E":
                    endPos = (x, y)
        return dataMap, startPos, endPos


def findPaths(dataMap, shortestPaths, endPos, currentPos):
    if currentPos == endPos:
        shortestPaths[endPos] = 0
        return 0

    shortestPaths[currentPos] = -1
    deltas = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    x, y = currentPos[0], currentPos[1]

    paths = []

    for d in deltas:
        nX, nY = x + d[0], y + d[1]
        newPos = (nX, nY)
        if not newPos in shortestPaths:
            if dataMap[newPos] != False and dataMap[newPos] != "#":
                tmpAns = findPaths(dataMap, shortestPaths, endPos, newPos) + 1
                if tmpAns != -1:
                    paths.append(tmpAns)

    ans = min(paths) if len(paths) > 0 else -1
    shortestPaths[currentPos] = ans
    return ans


def part1(data):
    dataMap, startPos, endPos = data[0], data[1], data[2]
    shortestPaths = {}
    shortest = findPaths(dataMap, shortestPaths, endPos, startPos)

    saved = defaultdict(int)
    deltas = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    for originalPos in shortestPaths.keys():
        x, y = originalPos[0], originalPos[1]
        for d1 in deltas:
            nX, nY = x + d1[0], y + d1[1]
            newPos = (nX, nY)
            if dataMap[newPos] == "#":
                for d2 in deltas:
                    cheatPos = (nX + d2[0], nY + d2[1])
                    if cheatPos != originalPos and cheatPos in shortestPaths:
                        originalDistance = shortest - shortestPaths[cheatPos]
                        newDistance = shortest - shortestPaths[originalPos] + 2
                        if newDistance + 100 <= originalDistance:
                            saved[originalDistance - newDistance] += 1

    return sum(saved.values())


def part2(data):
    dataMap, startPos, endPos = data[0], data[1], data[2]
    shortestPaths = {}
    shortest = findPaths(dataMap, shortestPaths, endPos, startPos)

    saved = defaultdict(int)
    maxD = 20
    for originalPos in shortestPaths.keys():
        x, y = originalPos[0], originalPos[1]
        for dX in range(-maxD, maxD + 1):
            for dY in range(-maxD + abs(dX), maxD - abs(dX) + 1):
                newPos = (x + dX, y + dY)
                if newPos in shortestPaths:
                    originalDistance = shortest - shortestPaths[newPos]
                    newDistance = shortest - shortestPaths[originalPos] + abs(dX) + abs(dY)
                    if newDistance + 100 <= originalDistance:
                        saved[originalDistance - newDistance] += 1

    return sum(saved.values())


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

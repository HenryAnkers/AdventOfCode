import os
import time
from collections import defaultdict


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def findExit(data, numToSimulate, endPos):
    dataMap = defaultdict(lambda: "#")

    for i in range(0, endPos[0] + 1):
        for j in range(0, endPos[1] + 1):
            dataMap[(i, j)] = "."

    for i in range(numToSimulate + 1):
        dataMap[(int(data[i].split(",")[0]), int(data[i].split(",")[1]))] = "#"

    queue = [(0, 0)]
    deltas = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    steps = 0
    while len(queue) > 0:
        newQueue = []
        for pos in queue:
            if pos == endPos:
                return steps
            x, y = pos[0], pos[1]
            for d in deltas:
                nX, nY = x + d[0], y + d[1]
                if dataMap[(nX, nY)] != "#":
                    newQueue.append((nX, nY))
            dataMap[pos] = "#"
        queue = list(set(newQueue))
        steps += 1

    return -1


def part1(data):
    return findExit(data, 1024, (70, 70))


def part2(data):
    # for i in range(1024, len(data)):
    #     if part1(data, i) == -1:
    #         return data[i]
    high = len(data) - 1
    low = 1024
    result = -1

    while low <= high:
        mid = (high + low) // 2
        if findExit(data, mid, (70, 70)) == -1:
            result = mid
            high = mid - 1
        else:
            low = mid + 1

    return data[result] if result != -1 else -1


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInput(input_path)
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

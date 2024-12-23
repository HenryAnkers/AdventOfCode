import os
import time
from collections import defaultdict


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def part1(data):
    neighbours = defaultdict(lambda: set())
    for line in data:
        n1, n2 = line.split("-")[0], line.split("-")[1]
        neighbours[n1].add(n2)
        neighbours[n2].add(n1)

    allThrees = set()
    for key, value in neighbours.items():
        allConnections = list(value)
        for i, n1 in enumerate(allConnections[:-1]):
            for _, n2 in enumerate(allConnections[i + 1 :]):
                if n2 in neighbours[n1]:
                    c = sorted([key, n1, n2])
                    if c[0][0] == "t" or c[1][0] == "t" or c[2][0] == "t":
                        allThrees.add((c[0], c[1], c[2]))

    return len(allThrees)


def part2(data):
    neighbours = defaultdict(lambda: set())
    for line in data:
        n1, n2 = line.split("-")[0], line.split("-")[1]
        neighbours[n1].add(n2)
        neighbours[n2].add(n1)

    allConnected = set()
    for key, value in neighbours.items():
        allConnections = list(value)
        for i, n1 in enumerate(allConnections[:-1]):
            for _, n2 in enumerate(allConnections[i + 1 :]):
                if n2 in neighbours[n1]:
                    c = sorted([key, n1, n2])
                    allConnected.add((c[0], c[1], c[2]))

    newLargest = set()
    allComputers = neighbours.keys()
    largestConnections = [list(i) for i in list(allConnected)]
    while len(largestConnections) > 1:
        newLargest = set()
        for connectionList in largestConnections:
            connection = set(connectionList)
            for computer in allComputers:
                if not computer in connection:
                    canConnect = True
                    for c in connection:
                        if not c in neighbours[computer]:
                            canConnect = False
                            break
                    if canConnect:
                        c = sorted(connectionList + [computer])
                        newLargest.add(",".join(c))

        largestConnections = [i.split(",") for i in list(newLargest)]

    return newLargest.pop()


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

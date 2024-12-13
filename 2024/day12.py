import os
import time
from collections import defaultdict


def parseInputAsMap(file_path):
    with open(file_path, "r") as file:
        data = [line.strip() for line in file]
        dataMap = defaultdict(lambda: defaultdict(bool))
        for y, line in enumerate(data):
            for x, char in enumerate(line):
                dataMap[x][y] = char
        return dataMap


def processRegion(data, dType, x, y):
    if data[x][y] != dType:
        return 0, 0

    data[x][y] = dType + "P"
    area = 1
    perim = 0
    for d in [[0, 1], [0, -1], [1, 0], [-1, 0]]:
        tX, tY = x + d[0], y + d[1]
        if data[tX][tY] == dType:
            tempArea, tempPerim = processRegion(data, dType, tX, tY)
            area += tempArea
            perim += tempPerim
        elif data[tX][tY] != dType + "P":
            perim += 1

    return area, perim


def part1(data):
    # area is easy but how do we work out perimeter?
    # we need perimeter for every neighbour (up,down,left,right) which doesn't belong to the region
    totalPrice = 0

    for x in list(data.keys()):
        for y in list(data[x].keys()):
            if data[x][y] != False and len(data[x][y]) != 2:
                type = data[x][y]
                area, perim = processRegion(data, data[x][y], x, y)
                totalPrice += perim * area

    return totalPrice


def processRegionPart2(data, dType, x, y, perimCoords, areaCoords):
    if data[x][y] != dType:
        return 0, 0

    data[x][y] = dType + "P"
    areaCoords.add((x, y))
    for d in [[0, 1], [0, -1], [1, 0], [-1, 0]]:
        tX, tY = x + d[0], y + d[1]
        if data[tX][tY] == dType:
            tempArea, perimCoords = processRegionPart2(data, dType, tX, tY, perimCoords, areaCoords)
        elif data[tX][tY] != dType + "P":
            perimCoords[(x, y)] += [(tX, tY)]

    return areaCoords, perimCoords


def part2(data):
    # we probably need the perimeter coordinates now?
    # just travel around the list of perimeter coordinates, every time you change direction the price increases by one?
    # every time we change direction we're at a corner, so just count the number of corners? this (in hindsight obviously) works and it's a lot easier than actually trying to traverse the perimeter
    totalPrice = 0

    for x in list(data.keys()):
        for y in list(data[x].keys()):
            if data[x][y] != False and len(data[x][y]) != 2:
                areaCoords, perimCoords = processRegionPart2(data, data[x][y], x, y, defaultdict(list), set())
                processedType = data[x][y]
                corners = 0

                for a in areaCoords:
                    aX, aY = a
                    for i, p1 in enumerate(perimCoords[a]):
                        for _, p2 in enumerate(perimCoords[a][i + 1 :]):
                            x1, y1 = p1
                            x2, y2 = p2

                            if abs(x2 - x1) == 1 and abs(y2 - y1) == 1:
                                corners += 1

                    if data[aX + 1][aY] == processedType and data[aX][aY + 1] == processedType and data[aX + 1][aY + 1] != processedType:
                        corners += 1
                    if data[aX - 1][aY] == processedType and data[aX][aY + 1] == processedType and data[aX - 1][aY + 1] != processedType:
                        corners += 1
                    if data[aX - 1][aY] == processedType and data[aX][aY - 1] == processedType and data[aX - 1][aY - 1] != processedType:
                        corners += 1
                    if data[aX + 1][aY] == processedType and data[aX][aY - 1] == processedType and data[aX + 1][aY - 1] != processedType:
                        corners += 1

                totalPrice += len(areaCoords) * corners

    return totalPrice


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
input_data = parseInputAsMap(input_path)
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")

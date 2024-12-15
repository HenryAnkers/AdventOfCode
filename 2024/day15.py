import os
import time
from collections import defaultdict


def parseInputAsMap1(file_path):
    with open(file_path, "r") as file:
        fr = file.read()
        mapStr, instructions = fr.split("\n\n")[0], "".join(fr.split("\n\n")[1].split())
        startPos = (0, 0)
        data = [line.strip() for line in mapStr.split("\n")]
        dataMap = defaultdict(lambda: "#")
        for y, line in enumerate(data):
            for x, char in enumerate(line):
                dataMap[(x, y)] = char
                if char == "@":
                    startPos = (x, y)

        return dataMap, instructions, startPos


def printMap(robotMap):
    xCoords = [coord[0] for coord in robotMap.keys()]
    yCoords = [coord[1] for coord in robotMap.keys()]

    minX, maxX = min(xCoords), max(xCoords)
    minY, maxY = min(yCoords), max(yCoords)

    for y in range(minY, maxY + 1):
        line = ""
        for x in range(minX, maxX + 1):
            line += robotMap[(x, y)]
        print(line)


def part1(robotMap, moves, initialPos):
    # start at robots coordinates (x,y), direction D.
    # If there's no box in direction D, move the robot.
    # otherwise if there is a box, switch current coordinate to the box, repeat the process... return canMove = true/false to let the other items know if they can move as well
    answer = 0
    deltaMap = {"^": (0, -1), "v": (0, 1), ">": (1, 0), "<": (-1, 0)}
    robotPos = initialPos

    for move in moves:
        dX, dY = deltaMap[move][0], deltaMap[move][1]
        x, y = robotPos[0], robotPos[1]
        positionsToMove = [(x, y)]
        canMove = True
        while True:
            nX, nY = x + dX, y + dY
            if robotMap[(nX, nY)] == "#":
                canMove = False
                break
            elif robotMap[(nX, nY)] == "O":
                positionsToMove.append((nX, nY))
                x, y = nX, nY
            else:
                break

        if canMove:
            for i, p in enumerate(reversed(positionsToMove)):
                nX, nY = p[0] + dX, p[1] + dY
                if i == len(positionsToMove) - 1:
                    robotPos = (nX, nY)
                robotMap[(nX, nY)] = robotMap[p]
                robotMap[p] = "."

    for key in robotMap.keys():
        if robotMap[key] == "O":
            answer += 100 * key[1] + key[0]

    return answer


def parseInputAsMap2(file_path):
    with open(file_path, "r") as file:
        fr = file.read()
        mapStr, instructions = fr.split("\n\n")[0], "".join(fr.split("\n\n")[1].split())
        startPos = (0, 0)
        data = [line.strip() for line in mapStr.split("\n")]
        dataMap = defaultdict(lambda: "#")
        for y, line in enumerate(data):
            x = 0
            for _, char in enumerate(line):
                if char != "@":
                    dataMap[(x, y)] = char if char != "O" else "["
                    dataMap[(x + 1, y)] = char if char != "O" else "]"
                else:
                    dataMap[(x, y)] = char
                    dataMap[(x + 1, y)] = "."
                    startPos = (x, y)
                x += 2

        return dataMap, instructions, startPos


def part2(robotMap, moves, initialPos):
    # this is basically the same problem just slightly more annoying to code
    # is there a nice way of adapting part 1? probably not ffs
    answer = 0
    deltaMap = {"^": (0, -1), "v": (0, 1), ">": (1, 0), "<": (-1, 0)}
    robotPos = initialPos

    for move in moves:
        dX, dY = deltaMap[move][0], deltaMap[move][1]
        x, y = robotPos[0], robotPos[1]
        positionsToMove = [(x, y)]
        coordsInStep = set([(x, y)])
        canMove = True

        while len(coordsInStep) > 0 and canMove:
            nextStep = set()
            for coord in coordsInStep:
                x, y = coord[0], coord[1]
                nX, nY = x + dX, y + dY
                if robotMap[(nX, nY)] == "#":
                    canMove = False
                    break
                elif robotMap[(nX, nY)] == "[" and not (nX, nY) in nextStep:
                    nextStep.add((nX, nY))
                    if dY != 0:
                        nextStep.add((nX + 1, nY))
                elif robotMap[(nX, nY)] == "]" and not (nX, nY) in nextStep:
                    nextStep.add((nX, nY))
                    if dY != 0:
                        nextStep.add((nX - 1, nY))

            positionsToMove += list(nextStep)
            coordsInStep = nextStep

        if canMove:
            for i, p in enumerate(reversed(positionsToMove)):
                nX, nY = p[0] + dX, p[1] + dY
                if i == len(positionsToMove) - 1:
                    robotPos = (nX, nY)
                robotMap[(nX, nY)] = robotMap[p]
                robotMap[p] = "."

    for key in robotMap.keys():
        if robotMap[key] == "[":
            answer += 100 * key[1] + key[0]

    return answer


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
robotMap, moves, initialPos = parseInputAsMap1(input_path)
end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")

start_time = time.time()
solution1 = part1(robotMap, moves, initialPos)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
robotMap, moves, initialPos = parseInputAsMap2(input_path)
solution2 = part2(robotMap, moves, initialPos)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")
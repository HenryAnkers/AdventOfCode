import os
import time
from collections import defaultdict

def parseInput(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]

def part1(data):
    gX, gY = 0,0        #guardX, guardY
    dataMap = defaultdict(lambda: defaultdict(lambda: False))
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            dataMap[x][y] = char
            if char == "^":
                gX, gY = x,y
                dataMap[x][y] = "."

    directions = [[0,-1], [1,0], [0,1], [-1,0]]
    currentDirectionIndex = 0
    count = 0
    while True:
        if dataMap[gX][gY] == ".":
            count += 1 
            dataMap[gX][gY] = "X"
        
        d = directions[currentDirectionIndex % 4]
        newX, newY = gX + d[0], gY + d[1]
        if not dataMap[newX][newY]:
            return count
        if dataMap[newX][newY] == "#":
            currentDirectionIndex += 1
        else:
            gX, gY = newX, newY
 
def vStr(x, y, dX, dY):
    return f"{x} - {y} - {dX} - {dY}"

def part2(data):
    gX, gY = 0,0
    dataMap = defaultdict(lambda: defaultdict(lambda: False))
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            dataMap[x][y] = char
            if char == "^":
                gX, gY = x,y

    directions = [[0,-1], [1,0], [0,1], [-1,0]]
    currentDirectionIndex = 0
    potentialLoopCount = 0

    while True:
        d = directions[currentDirectionIndex]
        newX, newY = gX + d[0], gY + d[1]
        if not dataMap[newX][newY]:
            return potentialLoopCount
        if dataMap[newX][newY] == "#":
            currentDirectionIndex = (currentDirectionIndex + 1) % 4
        else:
            if dataMap[newX][newY] != "^":
                # check what happens if we placed a barrier at newX, newY
                dataMap[newX][newY] = "#"
                tempX, tempY = gX, gY
                tempDirectionIndex = currentDirectionIndex
                tempFoundVectors = set()

                while True:
                    pDx, pDy = directions[tempDirectionIndex]
                    nextX, nextY = tempX + pDx, tempY + pDy

                    if not dataMap[nextX][nextY]:
                        break
                    if dataMap[nextX][nextY] == "#":
                        tempDirectionIndex = (tempDirectionIndex + 1) % 4
                    else:
                        tempX, tempY = nextX, nextY
                        tempVec =(tempX << 20) + (tempY << 10) + tempDirectionIndex
                        if tempVec in tempFoundVectors:
                            potentialLoopCount += 1
                            break
                        else:
                            tempFoundVectors.add(tempVec)
                
                dataMap[newX][newY] = "^"
            gX, gY = newX, newY


day = os.path.basename(__file__).split('.')[0].replace('day', '')
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
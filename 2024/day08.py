import os
import time
from collections import defaultdict

def parseInput(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]
    
def calculateAntinodeCoords(coord1, coord2):
    x1, y1, x2, y2 = coord1[0], coord1[1], coord2[0], coord2[1]
    dx, dy = (x2 - x1), (y2 - y1)
    a1x, a1y = x1 - dx, y1 - dy
    a2x, a2y = x2 + dx, y2 + dy

    return (a1x, a1y), (a2x, a2y)

def part1(data):
    dataMap = defaultdict(bool)
    antennas = defaultdict(list)
    antinodes = defaultdict(bool)
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            if char != ".":
                antennas[char].append((x,y))
            dataMap[(x,y)] = char
    
    for antListType in antennas.keys():
        antList = antennas[antListType]
        for i, coord1 in enumerate(antList):
            for _, coord2 in enumerate(antList[i+1:]):
                newAntinodes = calculateAntinodeCoords(coord1, coord2)
                for a in newAntinodes:
                    if dataMap[a]:
                        antinodes[a] = True

    return len(antinodes.keys())


def part2(data):
    dataMap = defaultdict(bool)
    antennas = defaultdict(list)
    antinodes = defaultdict(bool)
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            if char != ".":
                antennas[char].append((x,y))
            dataMap[(x,y)] = char
    
    for antListType in antennas.keys():
        antList = antennas[antListType]
        for i, coord1 in enumerate(antList):
            for _, coord2 in enumerate(antList[i+1:]):
                antinodes[coord1] = True
                antinodes[coord2] = True
                x1, y1, x2, y2 = coord1[0], coord1[1], coord2[0], coord2[1]
                dx, dy = (x2 - x1), (y2 - y1)

                tempAntinode = (x1 - dx, y1 - dy)         
                while dataMap[tempAntinode]:
                     xt, yt = tempAntinode[0], tempAntinode[1]
                     antinodes[tempAntinode] = True
                     tempAntinode = (xt - dx, yt - dy)         
                
                tempAntinode = (x2 + dx, y2 + dy)          
                while dataMap[tempAntinode]:
                    xt, yt = tempAntinode[0], tempAntinode[1]
                    antinodes[tempAntinode] = True
                    tempAntinode = (xt + dx, yt + dy)           


    return len(antinodes.keys())  


day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInput(input_path)
end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")
#print(input_data)

start_time = time.time()
solution1 = part1(input_data)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")